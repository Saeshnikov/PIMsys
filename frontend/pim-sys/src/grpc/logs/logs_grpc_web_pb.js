/**
 * @fileoverview gRPC-Web generated client stub for logs
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.21.12
// source: logs.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.logs = require('./logs_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.logs.LogsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.logs.LogsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.logs.GetLogsRequest,
 *   !proto.logs.GetLogsResponse>}
 */
const methodDescriptor_Logs_GetLogs = new grpc.web.MethodDescriptor(
  '/logs.Logs/GetLogs',
  grpc.web.MethodType.UNARY,
  proto.logs.GetLogsRequest,
  proto.logs.GetLogsResponse,
  /**
   * @param {!proto.logs.GetLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.logs.GetLogsResponse.deserializeBinary
);


/**
 * @param {!proto.logs.GetLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.logs.GetLogsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.logs.GetLogsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.logs.LogsClient.prototype.getLogs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/logs.Logs/GetLogs',
      request,
      metadata || {},
      methodDescriptor_Logs_GetLogs,
      callback);
};


/**
 * @param {!proto.logs.GetLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.logs.GetLogsResponse>}
 *     Promise that resolves to the response
 */
proto.logs.LogsPromiseClient.prototype.getLogs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/logs.Logs/GetLogs',
      request,
      metadata || {},
      methodDescriptor_Logs_GetLogs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.logs.GetGraphRequest,
 *   !proto.logs.GetGraphResponse>}
 */
const methodDescriptor_Logs_GetGraph = new grpc.web.MethodDescriptor(
  '/logs.Logs/GetGraph',
  grpc.web.MethodType.UNARY,
  proto.logs.GetGraphRequest,
  proto.logs.GetGraphResponse,
  /**
   * @param {!proto.logs.GetGraphRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.logs.GetGraphResponse.deserializeBinary
);


/**
 * @param {!proto.logs.GetGraphRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.logs.GetGraphResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.logs.GetGraphResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.logs.LogsClient.prototype.getGraph =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/logs.Logs/GetGraph',
      request,
      metadata || {},
      methodDescriptor_Logs_GetGraph,
      callback);
};


/**
 * @param {!proto.logs.GetGraphRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.logs.GetGraphResponse>}
 *     Promise that resolves to the response
 */
proto.logs.LogsPromiseClient.prototype.getGraph =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/logs.Logs/GetGraph',
      request,
      metadata || {},
      methodDescriptor_Logs_GetGraph);
};


module.exports = proto.logs;

