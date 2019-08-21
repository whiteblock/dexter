/**
 * @fileoverview gRPC-Web generated client stub for dexter
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.dexter = require('./data_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.dexter.DataClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.dexter.DataPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dexter.ExchangesRequest,
 *   !proto.dexter.ExchangesResponse>}
 */
const methodInfo_Data_SupportedExchanges = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dexter.ExchangesResponse,
  /** @param {!proto.dexter.ExchangesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.dexter.ExchangesResponse.deserializeBinary
);


/**
 * @param {!proto.dexter.ExchangesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dexter.ExchangesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dexter.ExchangesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dexter.DataClient.prototype.supportedExchanges =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dexter.Data/SupportedExchanges',
      request,
      metadata || {},
      methodInfo_Data_SupportedExchanges,
      callback);
};


/**
 * @param {!proto.dexter.ExchangesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dexter.ExchangesResponse>}
 *     A native promise that resolves to the response
 */
proto.dexter.DataPromiseClient.prototype.supportedExchanges =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dexter.Data/SupportedExchanges',
      request,
      metadata || {},
      methodInfo_Data_SupportedExchanges);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dexter.MarketsRequest,
 *   !proto.dexter.MarketsResponse>}
 */
const methodInfo_Data_SupportedMarkets = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dexter.MarketsResponse,
  /** @param {!proto.dexter.MarketsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.dexter.MarketsResponse.deserializeBinary
);


/**
 * @param {!proto.dexter.MarketsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dexter.MarketsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dexter.MarketsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dexter.DataClient.prototype.supportedMarkets =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dexter.Data/SupportedMarkets',
      request,
      metadata || {},
      methodInfo_Data_SupportedMarkets,
      callback);
};


/**
 * @param {!proto.dexter.MarketsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dexter.MarketsResponse>}
 *     A native promise that resolves to the response
 */
proto.dexter.DataPromiseClient.prototype.supportedMarkets =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dexter.Data/SupportedMarkets',
      request,
      metadata || {},
      methodInfo_Data_SupportedMarkets);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dexter.CandlesRequest,
 *   !proto.dexter.CandlesResponse>}
 */
const methodInfo_Data_GetCandles = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dexter.CandlesResponse,
  /** @param {!proto.dexter.CandlesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.dexter.CandlesResponse.deserializeBinary
);


/**
 * @param {!proto.dexter.CandlesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dexter.CandlesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dexter.CandlesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dexter.DataClient.prototype.getCandles =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dexter.Data/GetCandles',
      request,
      metadata || {},
      methodInfo_Data_GetCandles,
      callback);
};


/**
 * @param {!proto.dexter.CandlesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dexter.CandlesResponse>}
 *     A native promise that resolves to the response
 */
proto.dexter.DataPromiseClient.prototype.getCandles =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dexter.Data/GetCandles',
      request,
      metadata || {},
      methodInfo_Data_GetCandles);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dexter.CandlesRequest,
 *   !proto.dexter.Candle>}
 */
const methodInfo_Data_StreamCandles = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dexter.Candle,
  /** @param {!proto.dexter.CandlesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.dexter.Candle.deserializeBinary
);


/**
 * @param {!proto.dexter.CandlesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.dexter.Candle>}
 *     The XHR Node Readable Stream
 */
proto.dexter.DataClient.prototype.streamCandles =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/dexter.Data/StreamCandles',
      request,
      metadata || {},
      methodInfo_Data_StreamCandles);
};


/**
 * @param {!proto.dexter.CandlesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.dexter.Candle>}
 *     The XHR Node Readable Stream
 */
proto.dexter.DataPromiseClient.prototype.streamCandles =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/dexter.Data/StreamCandles',
      request,
      metadata || {},
      methodInfo_Data_StreamCandles);
};


module.exports = proto.dexter;

