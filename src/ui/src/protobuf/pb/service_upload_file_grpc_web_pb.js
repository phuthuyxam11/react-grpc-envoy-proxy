/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: service_upload_file.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pb = require('./service_upload_file_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.HrTimeSheetServiceClient =
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
proto.pb.HrTimeSheetServicePromiseClient =
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
 *   !proto.pb.FileUploadRequest,
 *   !proto.pb.FileUploadResponse>}
 */
const methodDescriptor_HrTimeSheetService_UploadFile = new grpc.web.MethodDescriptor(
  '/pb.HrTimeSheetService/UploadFile',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.pb.FileUploadRequest,
  proto.pb.FileUploadResponse,
  /**
   * @param {!proto.pb.FileUploadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.FileUploadResponse.deserializeBinary
);


/**
 * @param {!proto.pb.FileUploadRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.pb.FileUploadResponse>}
 *     The XHR Node Readable Stream
 */
proto.pb.HrTimeSheetServiceClient.prototype.uploadFile =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/pb.HrTimeSheetService/UploadFile',
      request,
      metadata || {},
      methodDescriptor_HrTimeSheetService_UploadFile);
};


/**
 * @param {!proto.pb.FileUploadRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.pb.FileUploadResponse>}
 *     The XHR Node Readable Stream
 */
proto.pb.HrTimeSheetServicePromiseClient.prototype.uploadFile =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/pb.HrTimeSheetService/UploadFile',
      request,
      metadata || {},
      methodDescriptor_HrTimeSheetService_UploadFile);
};


module.exports = proto.pb;
