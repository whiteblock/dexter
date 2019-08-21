import grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';
import dexterData from './index';
import GrpcBoom from 'grpc-boom';
import pino from 'pino';

const logger = pino();

const PROTO_PATH = `${__dirname}/../proto/data.proto`;
const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
    arrays: true,
  },
);
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

function unsupported(call: any, cb: any) {
  cb(GrpcBoom.unimplemented("This method has not been implemented yet"))
}

/**
 * @param call
 * @return metadata
 */
function copyMetadata(call: any) {
  const metadata = call.metadata.getMap();
  const responseMetadata = new grpc.Metadata();
  for (const key in metadata) {
    responseMetadata.set(key, metadata[key]);
  }
  return responseMetadata;
}

function supportedExchanges(call: any, cb: any) {
  logger.info({ method: 'SupportedExchanges' });
  const exchanges = dexterData.supportedExchanges();
  cb(null, { exchanges }, copyMetadata(call));
}

async function supportedMarkets(call: any, cb: any) {
  const {exchange} = call.request;
  logger.info({ method: 'SupportedMarkets', exchange });
  const markets = await dexterData.supportedMarkets(exchange);
  cb(null, { markets }, copyMetadata(call));
}

async function getCandles(call: any, cb: any) {
  const {exchange, market, timeframe, since, limit} = call.request;
  logger.info({ method: 'GetCandles', exchange, market, timeframe, since, limit });
  const candles = await dexterData.getCandles(exchange, market, timeframe, since, limit);
  if (candles && candles.length > 0) {
    cb(null, { candles: candles.map((c:any) => { return { timestamp: c[0], o: c[1], h: c[2], l: c[3], c: c[4], v: c[5] }}) }, copyMetadata(call));
  } else {
    cb(GrpcBoom.unavailable(`No candle data for ${exchange}/${market}/${timeframe}`));
  }
}

async function streamCandles(call: any) {
  const {exchange, market, timeframe} = call.request;
  logger.info({ method: 'StreamCandles', exchange, market, timeframe });
  const candleEmitter = await dexterData.streamCandles(exchange, market, timeframe);
  candleEmitter.output.on('candle', (candle) => {
    const c = {
      timestamp: candle[0],
      o: candle[1],
      h: candle[2],
      l: candle[3],
      c: candle[4],
      v: candle[5],
    }
    call.write(c);
  })
  call.on('cancelled', () => {
    logger.info({ method: 'StreamCandles', cancelled: true, exchange, market, timeframe })
    candleEmitter.stop();
  })
  candleEmitter.start();
}

function getServer() {
  const server = new grpc.Server();
  const dd: any = (protoDescriptor.dexter as any).Data;
  const service = dd.service;
  server.addService(service, {
    supportedExchanges,
    supportedMarkets,
    getCandles,
    streamCandles,
  });
  return server;
}

function startServer(bind: string) {
  logger.info({ method: 'StartServer', bind })
  const server = getServer();
  server.bind(bind, grpc.ServerCredentials.createInsecure());
  server.start();

}

function getClient(bind: string) {
  return new (protoDescriptor.dexter as any).Data(bind, grpc.credentials.createInsecure());
}

export default {
  packageDefinition,
  protoDescriptor,
  unsupported,
  getServer,
  startServer,
  getClient,
};
