/**
 * @fileoverview gRPC-Web generated client stub for shop
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.21.12
// source: shop.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.shop = require('./shop_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.shop.ShopClient =
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
proto.shop.ShopPromiseClient =
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
 *   !proto.shop.ListShopsRequest,
 *   !proto.shop.ListShopsResponse>}
 */
const methodDescriptor_Shop_ListShops = new grpc.web.MethodDescriptor(
  '/shop.Shop/ListShops',
  grpc.web.MethodType.UNARY,
  proto.shop.ListShopsRequest,
  proto.shop.ListShopsResponse,
  /**
   * @param {!proto.shop.ListShopsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.shop.ListShopsResponse.deserializeBinary
);


/**
 * @param {!proto.shop.ListShopsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.shop.ListShopsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.shop.ListShopsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.shop.ShopClient.prototype.listShops =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/shop.Shop/ListShops',
      request,
      metadata || {},
      methodDescriptor_Shop_ListShops,
      callback);
};


/**
 * @param {!proto.shop.ListShopsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.shop.ListShopsResponse>}
 *     Promise that resolves to the response
 */
proto.shop.ShopPromiseClient.prototype.listShops =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/shop.Shop/ListShops',
      request,
      metadata || {},
      methodDescriptor_Shop_ListShops);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.shop.NewShopRequest,
 *   !proto.shop.NewShopResponse>}
 */
const methodDescriptor_Shop_NewShop = new grpc.web.MethodDescriptor(
  '/shop.Shop/NewShop',
  grpc.web.MethodType.UNARY,
  proto.shop.NewShopRequest,
  proto.shop.NewShopResponse,
  /**
   * @param {!proto.shop.NewShopRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.shop.NewShopResponse.deserializeBinary
);


/**
 * @param {!proto.shop.NewShopRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.shop.NewShopResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.shop.NewShopResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.shop.ShopClient.prototype.newShop =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/shop.Shop/NewShop',
      request,
      metadata || {},
      methodDescriptor_Shop_NewShop,
      callback);
};


/**
 * @param {!proto.shop.NewShopRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.shop.NewShopResponse>}
 *     Promise that resolves to the response
 */
proto.shop.ShopPromiseClient.prototype.newShop =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/shop.Shop/NewShop',
      request,
      metadata || {},
      methodDescriptor_Shop_NewShop);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.shop.AlterShopRequest,
 *   !proto.shop.AlterShopResponse>}
 */
const methodDescriptor_Shop_AlterShop = new grpc.web.MethodDescriptor(
  '/shop.Shop/AlterShop',
  grpc.web.MethodType.UNARY,
  proto.shop.AlterShopRequest,
  proto.shop.AlterShopResponse,
  /**
   * @param {!proto.shop.AlterShopRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.shop.AlterShopResponse.deserializeBinary
);


/**
 * @param {!proto.shop.AlterShopRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.shop.AlterShopResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.shop.AlterShopResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.shop.ShopClient.prototype.alterShop =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/shop.Shop/AlterShop',
      request,
      metadata || {},
      methodDescriptor_Shop_AlterShop,
      callback);
};


/**
 * @param {!proto.shop.AlterShopRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.shop.AlterShopResponse>}
 *     Promise that resolves to the response
 */
proto.shop.ShopPromiseClient.prototype.alterShop =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/shop.Shop/AlterShop',
      request,
      metadata || {},
      methodDescriptor_Shop_AlterShop);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.shop.DeleteShopRequest,
 *   !proto.shop.DeleteShopResponse>}
 */
const methodDescriptor_Shop_DeleteShop = new grpc.web.MethodDescriptor(
  '/shop.Shop/DeleteShop',
  grpc.web.MethodType.UNARY,
  proto.shop.DeleteShopRequest,
  proto.shop.DeleteShopResponse,
  /**
   * @param {!proto.shop.DeleteShopRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.shop.DeleteShopResponse.deserializeBinary
);


/**
 * @param {!proto.shop.DeleteShopRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.shop.DeleteShopResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.shop.DeleteShopResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.shop.ShopClient.prototype.deleteShop =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/shop.Shop/DeleteShop',
      request,
      metadata || {},
      methodDescriptor_Shop_DeleteShop,
      callback);
};


/**
 * @param {!proto.shop.DeleteShopRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.shop.DeleteShopResponse>}
 *     Promise that resolves to the response
 */
proto.shop.ShopPromiseClient.prototype.deleteShop =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/shop.Shop/DeleteShop',
      request,
      metadata || {},
      methodDescriptor_Shop_DeleteShop);
};


module.exports = proto.shop;

