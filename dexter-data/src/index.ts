import _ccxt from 'ccxt';
import { EventEmitter } from 'events';
import pino from 'pino';
import time from './time';

const logger = pino();
const ccxt: any = _ccxt; // a hack to make TypeScript shut up.

/*

  Requirements

  - Provide a way to ask for candlestick data from an exchange and market and timeframe.
    + If candlesticks of a certain timeframe are not available,
      try to derive them from smaller timeframes that *are* available.
  - Provide a way to start and stop continuous download of candlesticks for a market.
  - Keep metadata on what markets I have data for.
  - Provide a way to stream candlesticks out to consumers.
    + Optionally, clean the candlestick data if they ask.
  - Wrap all this in a gRPC service.

 */


interface PriceEmitterOptions { // or CandleEmitter
  exchange:        string;
  market:          string;
  updateInterval?: number;
}

class PriceEmitter {
  intervalId:     NodeJS.Timeout | undefined;
  updateInterval: number;
  em:             EventEmitter;
  ex:             any;
  exchange:       string;
  market:         string;
  subscribers:    Array<EventEmitter>;

  constructor(opts: PriceEmitterOptions) {
    this.updateInterval = opts.updateInterval || 10000;
    this.ex = new ccxt[opts.exchange]();
    this.exchange = opts.exchange;
    this.market = opts.market;
    const em = this.em = new EventEmitter();
    this.subscribers = [];
    em.on('price', (candles) => {
      // loop through every subscriber and send them the last 2 1m candles
      this.subscribers.forEach((sub) => {
        sub.emit('price', candles)
      })
    });
  }

  start() {
    if (this.intervalId) return;
    this.intervalId = setInterval(async () => {
      try {
        const candles = await this.ex.fetchOHLCV(this.market, '1m', undefined, 2);
        if (candles[1]) {
          this.em.emit('price', candles);
        }
      }
      catch (e) {
        console.warn(e);
      }
    }, this.updateInterval);
  }

  stop() {
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = undefined
    }
  }

  addSubscriber(subscriber: EventEmitter) {
    this.subscribers.push(subscriber)
  }

  // TODO
  removeSubscriber(subscriber: EventEmitter) {
  }
}

interface CandleEmitterOptions {
  exchange: string
  market: string
  timeframe: string
  closeOnly: boolean
}

class CandleEmitter {
  lastCandle: Array<number>
  input: EventEmitter
  output: EventEmitter
  exchange: string
  market: string
  timeframe: string
  closeOnly: boolean
  ex: any
  volumeAccumulator: number // TODO - keep track of volume
  priceHandler: any // TODO - replace with a function type

  constructor(opts: CandleEmitterOptions) {
    this.lastCandle        = []
    this.exchange          = opts.exchange
    this.market            = opts.market
    this.timeframe         = opts.timeframe
    this.closeOnly         = opts.closeOnly
    this.ex                = new ccxt[opts.exchange]()
    this.volumeAccumulator = 0
    this.output            = new EventEmitter()
    this.input             = new EventEmitter()
    this.priceHandler      = this.emitCandles.bind(this)
  }

  /**
   * Start emitting candles.
   */
  async start() {
    // initialize this.lastCandle with timeframe-appropriate data
    const candles = await this.ex.fetchOHLCV(this.market, this.timeframe, undefined, 2)
    this.lastCandle = candles[0]
    this.volumeAccumulator = candles[5]
    // generate candles as new price data comes in
    this.input.on('price', this.priceHandler)
  }

  /**
   * Stop emitting candles.
   */
  stop() {
    this.input.off('price', this.priceHandler);
  }

  /**
   * Compare the `lastCandle` with the current `candle` and return an updated candle in the same timeframe.
   * @param lastCandle  Previous candle data
   * @param candle      The most recent candle data
   * @returns           An updated candle
   */
  updateCandle(lastCandle: Array<number>, candle: Array<number>): Array<number> {
    // 0 timestamp
    // 1 open
    // 2 high
    // 3 low
    // 4 close
    // 5 volume
    const close = candle[4];
    const newCandle = [...lastCandle];
    if (lastCandle.length === 0) {
      newCandle[1] = newCandle[2] = newCandle[3] = newCandle[4] = close;
      newCandle[0] = time.timestampForTimeframe(this.timeframe, candle[0]);
      return newCandle;
    }
    if (newCandle[2] < close) {
      newCandle[2] = close
    }
    if (newCandle[3] > close) {
      newCandle[3] = close
    }
    newCandle[4] = close;
    const newVolume = lastCandle[5] + newCandle[5]
    newCandle[5] = newVolume
    return newCandle
  }

  /**
   * This takes 'price' messages emitted from a PriceEmitter and creates updated candles.
   * @param candles pair of candles from an exchange
   */
  emitCandles(candles: Array<Array<number>>) {
    // the reason I'm sending the last 2 candles is because when a new candle opens, I don't want to miss
    // the last bit of data
    if (candles.length === 0) {
      // no data
      return
    } else if (candles.length === 1) {
      // only one candle available.  brand new market?
      this.lastCandle = candles[0]
      this.lastCandle[0] = time.timestampForTimeframe(this.timeframe, candles[0][0]);
      this.output.emit('candle', this.lastCandle);
      return
    } else if (candles.length === 2) {
      // typical case
      const c0 = candles[0];
      const c1 = candles[1];
      const c1ts = time.dt(c1[0]);
      if (time.isTimeframeBoundary(this.timeframe, c1ts)) {
        if (this.lastCandle[4] != c0[4]) {
          const finishCandle = this.updateCandle(this.lastCandle, c0);
          this.output.emit('candle', finishCandle);
          this.lastCandle = c1;
          logger.info({ method: 'emitCandles', candle: this.lastCandle })
          this.output.emit('candle', this.lastCandle);
        }
        this.lastCandle = c1;
      } else {
        this.lastCandle = this.updateCandle(this.lastCandle, c1);
        logger.info({ method: 'emitCandles', candle: this.lastCandle })
        this.output.emit('candle', this.lastCandle);
      }
    }
  }
}

/**
 * Return a list of exchanges we can pull market data from.
 * @returns A list of exchanges we support.
 */
function supportedExchanges(): Array<string> {
  return ccxt.exchanges
    .map((name: string) => new ccxt[name]())
    .filter((ex: any) => ex.has['fetchOHLCV'])
    .map((ex: any) => ex.id)
    .toArray()
}

/**
 * Return a list of markets available on an exchange.
 * @param exchange An exchange name
 * @returns a list of markets
 */
async function supportedMarkets(exchange: string): Promise<Array<string>> {
  const ex: any = new ccxt[exchange]();
  const markets = await ex.loadMarkets();
  const symbols = Object.keys(markets);
  return symbols;
}

/**
 * Return a list of candlesticks for a given market.
 * @param exchange An exchange name
 * @param market A market
 * @param timeframe The duration of a candlestick
 * @param since A point in time in the past in ISO-8601 format.
 * @param limit The number greater than 0 of the number of candles to fetch.
 */
async function getCandles(exchange: string, market: string, timeframe: string, since: string, limit: number) {
  const ex: any = new ccxt[exchange]();
  const args: Array<any> = [ market, timeframe ];
  const timeframes: Array<string> = Object.keys(ex.timeframes)
  if (timeframes.includes(timeframe)) {
    // Timeframe is natively supported by the exchange
    if (since) {
      args.push(ex.parse8601(since))
    } else {
      args.push(undefined)
    }
    if (limit > 0) {
      args.push(Number(limit))
    }
    const candles: Array<Object> = await ex.fetchOHLCV(...args);
    return candles
  } else {
    // If the timeframe is not natively supported by the exchange, try to emulate it.
    const nativeTimeframes = Object.keys(ex.timeframes);
    const reverseMap       = time.reverseMapMinutesToTimeframes(nativeTimeframes);
    const commonMinutes    = time.highestCommonTimeframe(nativeTimeframes, timeframe);
    const commonTimeframe  = reverseMap[commonMinutes]
    if (commonTimeframe) {
      args[1] = commonTimeframe
      if (since) {
        args.push(ex.parse8601(since))
      } else {
        args.push(undefined)
      }
      if (limit > 0) {
        args.push(Number(limit))
      }
      const candles: Array<Object> = await ex.fetchOHLCV(...args);
      return time.emulateTimeframeCandles(timeframe, candles)
    }
  }
}

// PriceEmitter -[1m candles every interval]-> CandleAggregator -[new candle for given timeframe]-> 

const emitters: {[key:string]: PriceEmitter} = {};

/**
 * Return a `PriceEmitter` for the given `exchange` and `market`, constructing a new one if necessary.
 * @param exchange An exchange name for ccxt
 * @param market   A market in `exchange`
 * @param updateInterval Milliseconds between price fetches
 * @returns A `PriceEmitter` instance
 */
function getPriceEmitter(exchange: string, market: string, updateInterval: number = 10000): PriceEmitter {
  const key = `${exchange}|${market}`;
  if (emitters[key]) {
    return emitters[key];
  } else {
    const em = new PriceEmitter({ exchange, market, updateInterval });
    emitters[key] = em;
    return em;
  }
}

/**
 * Return a `CandleEmitter` for the given `exchange`, `market`, and `timeframe`
 * @param exchange An exchange name for ccxt
 * @param market A market in `exchange`
 * @param timeframe The duration of a candlestick
 * @returns A `CandleEmitter` instance
 */
async function streamCandles(exchange: string, market: string, timeframe: string) {
  const priceEm = getPriceEmitter(exchange, market);
  // TODO need to initialize it with the current candle.
  const candleEmitter = new CandleEmitter({ exchange, market, timeframe, closeOnly: false })
  priceEm.addSubscriber(candleEmitter.input);
  priceEm.start();
  return candleEmitter;
}

/**
 * Start saving candlesticks to a database.
 * @param exchange An exchange name for ccxt
 * @param market A market in `exchange` 
 * @param timeframe The duration of a candlestick
 */
function archiveMarket(exchange: string, market: string, timeframe: string) {
  // TODO - Implement this method.
}

export default {
  //Bookshelf,
  //Knex,
  //ccxt,
  //Lazy,
  PriceEmitter,
  CandleEmitter,
  supportedExchanges,
  supportedMarkets,
  getCandles,
  getPriceEmitter,
  streamCandles,
  archiveMarket,
};
