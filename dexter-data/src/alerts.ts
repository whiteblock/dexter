import grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';

const PROTO_PATH = `${__dirname}/../proto/alerts.proto`;
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

function getClient(bind: string) {
  return new (protoDescriptor.dexter as any).Alerts(bind, grpc.credentials.createInsecure());
}

export default {
  getClient,
};
