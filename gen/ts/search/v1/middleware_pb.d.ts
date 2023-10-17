// @generated by protoc-gen-es v1.3.3
// @generated from file search/v1/middleware.proto (package search.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { GlobalContract, LocalContract } from "./contracts_pb.js";
import type { RemoteParticipant } from "./broker_pb.js";

/**
 * @generated from message search.v1.AppSendResponse
 */
export declare class AppSendResponse extends Message<AppSendResponse> {
  /**
   * @generated from field: search.v1.AppSendResponse.Result result = 1;
   */
  result: AppSendResponse_Result;

  constructor(data?: PartialMessage<AppSendResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.AppSendResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppSendResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppSendResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppSendResponse;

  static equals(a: AppSendResponse | PlainMessage<AppSendResponse> | undefined, b: AppSendResponse | PlainMessage<AppSendResponse> | undefined): boolean;
}

/**
 * @generated from enum search.v1.AppSendResponse.Result
 */
export declare enum AppSendResponse_Result {
  /**
   * @generated from enum value: RESULT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESULT_OK = 1;
   */
  OK = 1,

  /**
   * @generated from enum value: RESULT_ERR = 2;
   */
  ERR = 2,
}

/**
 * @generated from message search.v1.AppRecvRequest
 */
export declare class AppRecvRequest extends Message<AppRecvRequest> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * @generated from field: string participant = 2;
   */
  participant: string;

  constructor(data?: PartialMessage<AppRecvRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.AppRecvRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppRecvRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppRecvRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppRecvRequest;

  static equals(a: AppRecvRequest | PlainMessage<AppRecvRequest> | undefined, b: AppRecvRequest | PlainMessage<AppRecvRequest> | undefined): boolean;
}

/**
 * @generated from message search.v1.RegisterChannelRequest
 */
export declare class RegisterChannelRequest extends Message<RegisterChannelRequest> {
  /**
   * @generated from field: search.v1.GlobalContract requirements_contract = 1;
   */
  requirementsContract?: GlobalContract;

  /**
   * Mapping of participants that we don't want brokered.
   *
   * @generated from field: map<string, search.v1.RemoteParticipant> preset_participants = 2;
   */
  presetParticipants: { [key: string]: RemoteParticipant };

  constructor(data?: PartialMessage<RegisterChannelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.RegisterChannelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterChannelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterChannelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterChannelRequest;

  static equals(a: RegisterChannelRequest | PlainMessage<RegisterChannelRequest> | undefined, b: RegisterChannelRequest | PlainMessage<RegisterChannelRequest> | undefined): boolean;
}

/**
 * @generated from message search.v1.RegisterChannelResponse
 */
export declare class RegisterChannelResponse extends Message<RegisterChannelResponse> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  constructor(data?: PartialMessage<RegisterChannelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.RegisterChannelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterChannelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterChannelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterChannelResponse;

  static equals(a: RegisterChannelResponse | PlainMessage<RegisterChannelResponse> | undefined, b: RegisterChannelResponse | PlainMessage<RegisterChannelResponse> | undefined): boolean;
}

/**
 * @generated from message search.v1.RegisterAppRequest
 */
export declare class RegisterAppRequest extends Message<RegisterAppRequest> {
  /**
   * @generated from field: search.v1.LocalContract provider_contract = 1;
   */
  providerContract?: LocalContract;

  constructor(data?: PartialMessage<RegisterAppRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.RegisterAppRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterAppRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterAppRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterAppRequest;

  static equals(a: RegisterAppRequest | PlainMessage<RegisterAppRequest> | undefined, b: RegisterAppRequest | PlainMessage<RegisterAppRequest> | undefined): boolean;
}

/**
 * whenever a new channel that involves this app is started, the middleware needs to notify the local app
 *
 * @generated from message search.v1.RegisterAppResponse
 */
export declare class RegisterAppResponse extends Message<RegisterAppResponse> {
  /**
   * @generated from oneof search.v1.RegisterAppResponse.ack_or_new
   */
  ackOrNew: {
    /**
     * @generated from field: string app_id = 1;
     */
    value: string;
    case: "appId";
  } | {
    /**
     * @generated from field: search.v1.InitChannelNotification notification = 2;
     */
    value: InitChannelNotification;
    case: "notification";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<RegisterAppResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.RegisterAppResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterAppResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterAppResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterAppResponse;

  static equals(a: RegisterAppResponse | PlainMessage<RegisterAppResponse> | undefined, b: RegisterAppResponse | PlainMessage<RegisterAppResponse> | undefined): boolean;
}

/**
 * this is what a registered app receives whenever a new channel is initiated for that app
 * the app has to communicate with the middleware using UseChannel with this new channel_id
 *
 * @generated from message search.v1.InitChannelNotification
 */
export declare class InitChannelNotification extends Message<InitChannelNotification> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  constructor(data?: PartialMessage<InitChannelNotification>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.InitChannelNotification";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InitChannelNotification;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InitChannelNotification;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InitChannelNotification;

  static equals(a: InitChannelNotification | PlainMessage<InitChannelNotification> | undefined, b: InitChannelNotification | PlainMessage<InitChannelNotification> | undefined): boolean;
}

/**
 * This is something that is sent by the Broker to providers to notify that a new channel is starting
 *
 * @generated from message search.v1.InitChannelRequest
 */
export declare class InitChannelRequest extends Message<InitChannelRequest> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * which app behind the middleware is being notified
   *
   * @generated from field: string app_id = 2;
   */
  appId: string;

  /**
   * int32 seq = 4; // sequence number (used because we may need multiple rounds until all participants are ready)
   *
   * @generated from field: map<string, search.v1.RemoteParticipant> participants = 3;
   */
  participants: { [key: string]: RemoteParticipant };

  constructor(data?: PartialMessage<InitChannelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.InitChannelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InitChannelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InitChannelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InitChannelRequest;

  static equals(a: InitChannelRequest | PlainMessage<InitChannelRequest> | undefined, b: InitChannelRequest | PlainMessage<InitChannelRequest> | undefined): boolean;
}

/**
 * @generated from message search.v1.InitChannelResponse
 */
export declare class InitChannelResponse extends Message<InitChannelResponse> {
  /**
   * @generated from field: search.v1.InitChannelResponse.Result result = 1;
   */
  result: InitChannelResponse_Result;

  constructor(data?: PartialMessage<InitChannelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.InitChannelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InitChannelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InitChannelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InitChannelResponse;

  static equals(a: InitChannelResponse | PlainMessage<InitChannelResponse> | undefined, b: InitChannelResponse | PlainMessage<InitChannelResponse> | undefined): boolean;
}

/**
 * @generated from enum search.v1.InitChannelResponse.Result
 */
export declare enum InitChannelResponse_Result {
  /**
   * @generated from enum value: RESULT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESULT_ACK = 1;
   */
  ACK = 1,

  /**
   * @generated from enum value: RESULT_ERR = 2;
   */
  ERR = 2,
}

/**
 * @generated from message search.v1.StartChannelRequest
 */
export declare class StartChannelRequest extends Message<StartChannelRequest> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * @generated from field: string app_id = 2;
   */
  appId: string;

  constructor(data?: PartialMessage<StartChannelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.StartChannelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartChannelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartChannelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartChannelRequest;

  static equals(a: StartChannelRequest | PlainMessage<StartChannelRequest> | undefined, b: StartChannelRequest | PlainMessage<StartChannelRequest> | undefined): boolean;
}

/**
 * @generated from message search.v1.StartChannelResponse
 */
export declare class StartChannelResponse extends Message<StartChannelResponse> {
  /**
   * @generated from field: search.v1.StartChannelResponse.Result result = 1;
   */
  result: StartChannelResponse_Result;

  constructor(data?: PartialMessage<StartChannelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.StartChannelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartChannelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartChannelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartChannelResponse;

  static equals(a: StartChannelResponse | PlainMessage<StartChannelResponse> | undefined, b: StartChannelResponse | PlainMessage<StartChannelResponse> | undefined): boolean;
}

/**
 * @generated from enum search.v1.StartChannelResponse.Result
 */
export declare enum StartChannelResponse_Result {
  /**
   * @generated from enum value: RESULT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESULT_ACK = 1;
   */
  ACK = 1,

  /**
   * @generated from enum value: RESULT_ERR = 2;
   */
  ERR = 2,
}

/**
 * @generated from message search.v1.CloseChannelRequest
 */
export declare class CloseChannelRequest extends Message<CloseChannelRequest> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  constructor(data?: PartialMessage<CloseChannelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.CloseChannelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloseChannelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloseChannelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloseChannelRequest;

  static equals(a: CloseChannelRequest | PlainMessage<CloseChannelRequest> | undefined, b: CloseChannelRequest | PlainMessage<CloseChannelRequest> | undefined): boolean;
}

/**
 * @generated from message search.v1.CloseChannelResponse
 */
export declare class CloseChannelResponse extends Message<CloseChannelResponse> {
  /**
   * @generated from field: search.v1.CloseChannelResponse.Result result = 1;
   */
  result: CloseChannelResponse_Result;

  /**
   * @generated from field: string error_message = 2;
   */
  errorMessage: string;

  /**
   * @generated from field: repeated string participants_with_pending_inbound = 3;
   */
  participantsWithPendingInbound: string[];

  constructor(data?: PartialMessage<CloseChannelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.CloseChannelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CloseChannelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CloseChannelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CloseChannelResponse;

  static equals(a: CloseChannelResponse | PlainMessage<CloseChannelResponse> | undefined, b: CloseChannelResponse | PlainMessage<CloseChannelResponse> | undefined): boolean;
}

/**
 * @generated from enum search.v1.CloseChannelResponse.Result
 */
export declare enum CloseChannelResponse_Result {
  /**
   * @generated from enum value: RESULT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESULT_CLOSED = 1;
   */
  CLOSED = 1,

  /**
   * @generated from enum value: RESULT_PENDING_INBOUND = 2;
   */
  PENDING_INBOUND = 2,

  /**
   * @generated from enum value: RESULT_PENDING_OUTBOUND = 3;
   */
  PENDING_OUTBOUND = 3,
}

/**
 * @generated from message search.v1.MessageExchangeResponse
 */
export declare class MessageExchangeResponse extends Message<MessageExchangeResponse> {
  /**
   * @generated from field: search.v1.MessageExchangeResponse.Result result = 1;
   */
  result: MessageExchangeResponse_Result;

  /**
   * @generated from field: string error_message = 2;
   */
  errorMessage: string;

  constructor(data?: PartialMessage<MessageExchangeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.MessageExchangeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageExchangeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageExchangeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageExchangeResponse;

  static equals(a: MessageExchangeResponse | PlainMessage<MessageExchangeResponse> | undefined, b: MessageExchangeResponse | PlainMessage<MessageExchangeResponse> | undefined): boolean;
}

/**
 * @generated from enum search.v1.MessageExchangeResponse.Result
 */
export declare enum MessageExchangeResponse_Result {
  /**
   * @generated from enum value: RESULT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: RESULT_OK = 1;
   */
  OK = 1,

  /**
   * @generated from enum value: RESULT_ERROR = 2;
   */
  ERROR = 2,
}

