import { DateTime, Interval } from 'luxon';

/**
 * Create a DateTime instance from milliseconds.
 * @param ms A timestamp in milliseconds
 * @returns A DateTime instance
 */
function dt(ms: number): DateTime {
  return DateTime.fromMillis(ms, { zone: 'UTC' });
}

/**
 * Return true if the given `time` is on an timeframe boundary.
 *
 * @param timeframe  InfluxDB duration notation, eg. 1m, 5m, 1d
 * @param time       A time
 * @returns          `true` if the timestamp is on an timeframe boundary.`
 */
function isTimeframeBoundary(
  timeframe: string,
  time: DateTime) : boolean {

  const match  = timeframe.match(/(\d+)(\w+)/);
  if (!match) {
    return false;
  }
  const nu = match[1];
  const unit = match[2];
  const n = Math.min(parseInt(nu, 10));
  const dayOfYear = Math.floor(
    Interval.fromDateTimes(DateTime.utc(time.year, 1, 1), time).length() + 1);

  switch (unit) {
    case 'm':
      if (time.minute % n === 0) return true;
      break;
    case 'h':
      if (time.hour % n === 0 && time.minute === 0) return true;
      break;
    case 'd':
      if (time.minute === 0 && time.hour === 0 && dayOfYear % n === 0) return true;
      break;
  }
  return false;
}

/**
 * Convert an timeframe into minutes.
 * @param   timeframe  A duration in InfluxDB notation
 * @returns            The duration of the timeframe in milliseconds
 */
function timeframeToMinutes(timeframe: string): number {
  const match  = timeframe.match(/(\d+)(\w+)/);
  if (!match) {
    throw new Error(`Invalid timeframe: '${timeframe}'`);
  }
  const nu = match[1];
  const unit = match[2];
  const n = Math.min(parseInt(nu, 10));
  // const now = DateTime.local();
  // const dayOfYear = Math.floor(
  //   Timeframe.fromDateTimes(DateTime.local(now.year, 1, 1), now).length() + 1);
  switch (unit) {
    case 'm':
      return n;
    case 'h':
      return 60 * n;
    case 'd':
      return 24 * 60 * n;
  }
  throw new Error(`Unsupported timeframe: '${timeframe}'`);
}

/**
 * Given an timeframe and a timestamp, adjust the timestamp so that it fits inside the timeframe.
 * @param timeframe A duration in InfluxDB notation
 * @param ms        A timestamp in milliseconds
 * @returns         An adjusted timestamp in milliseconds that fits inside `timeframe`
 */
function timestampForTimeframe(timeframe: string, ms: number): number {
  const ints = timeframeToMinutes(timeframe) * 60 * 1000;
  const diff = ms % ints;
  return ms - diff;
}

/**
 * Normalize a list of timeframes to minutes
 * @param timeframes An array of timeframes supported by an exchange
 * @returns          A list of minutes
 */
function minutesFromTimeframes(timeframes: Array<string>): Array<number> {
  return timeframes.map(timeframeToMinutes);
}

/**
 * Translate raw minutes back to timeframe notation
 * @param timeframes An array of timeframes supported by an exchange
 * @returns          A hashmap with raw minutes as keys and timeframes as values
 */
function reverseMapMinutesToTimeframes(timeframes: Array<string>): { [minutes: number]: string } {
  return timeframes.reduce((m: { [minutes: number]: string }, a) => {
    try {
      const minutes = timeframeToMinutes(a);
      m[minutes] = a;
      return m;
    } catch {
      return m;
    }
  }, {});
}


/**
 * Return the highest native timeframe that is evenly divisible into the given timeframe.
 * @param nativeTimeframes An array of timeframes supported by an exchange.
 * @param timeframe        The timeframe we want to emulate later
 * @returns                The highest common timeframe
 */
function highestCommonTimeframe(nativeTimeframes: Array<string>, timeframe: string): number {
  const minutes = minutesFromTimeframes(nativeTimeframes).sort((a, b) => b - a);
  const base = timeframeToMinutes(timeframe);
  // console.log({ minutes });
  const common = minutes.find((t) => {
    // console.log(`base(${base}) % t(${t}) == ${ base % 5 }`);
    return (base % t) === 0;
  });
  if (common) {
    return common;
  }
  throw(new Error(`Common timeframe not found for ${timeframe}`));
}

/**
 * Aggregate candles into an arbitrary timeframe
 * @param timeframe   The desired timeframe
 * @param candles     An array of candles in a timeframe that's smaller and even divisible into `timeframe`
 * @returns           An array of candles aggregated into `timeframe`
 */
function emulateTimeframeCandles(timeframe: string, candles: any): any {
  const emulatedCandles: any = [];
  candles.reduce((m: any, c: any) => {
    // if m is empty, set the right timestamp on c and let it be the first candle in m.
    if (m.length === 0) {
      c[0] = timestampForTimeframe(timeframe, c[0]);
      m.push(c);
    } else {
      if (isTimeframeBoundary(timeframe, c[0])) {
        // if candle c is on a timeframe boundary, start a new candle and push it on to m.
        m.push(c);
      } else {
        // else merge c into last candle.
        const lastCandle = m[m.length - 1];
        const newCandle = mergeCandle(lastCandle, c);
        m[m.length - 1] = newCandle;
      }
    }
    return m
  }, emulatedCandles);
  return emulatedCandles;
}

function mergeCandle(lastCandle: Array<number>, candle: Array<number>): Array<number> {
  const close = candle[4];
  const newCandle = [...lastCandle];
  if (newCandle[2] < close) {
    newCandle[2] = close;
  }
  if (newCandle[3] > close) {
    newCandle[3] = close;
  }
  newCandle[4] = close;
  const newVolume = lastCandle[5] + newCandle[5];
  newCandle[5] = newVolume;
  return newCandle;
}

export default {
  dt,
  isTimeframeBoundary,
  timeframeToMinutes,
  timestampForTimeframe,
  minutesFromTimeframes,
  reverseMapMinutesToTimeframes,
  highestCommonTimeframe,
  emulateTimeframeCandles,
  mergeCandle,
};
