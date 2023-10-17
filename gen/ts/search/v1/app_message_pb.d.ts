// @generated by protoc-gen-es v1.3.3
// @generated from file search/v1/app_message.proto (package search.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * This is what will be exchanged between middlewares
 *
 * @generated from message search.v1.MessageExchangeRequest
 */
export declare class MessageExchangeRequest extends Message<MessageExchangeRequest> {
  /**
   * We'll use UUIDv4. It's a global ID shared by all participants
   *
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * This is necessary because URLs don't univocally determine apps. There can be multiple applications
   * behind the same middleware (there is a 1:1 mapping between URLs and middlewares)
   *
   * appid de app emisora
   *
   * @generated from field: string sender_id = 2;
   */
  senderId: string;

  /**
   * appid de app receptora
   *
   * @generated from field: string recipient_id = 3;
   */
  recipientId: string;

  /**
   * @generated from field: search.v1.AppMessage content = 4;
   */
  content?: AppMessage;

  constructor(data?: PartialMessage<MessageExchangeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.MessageExchangeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageExchangeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageExchangeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageExchangeRequest;

  static equals(a: MessageExchangeRequest | PlainMessage<MessageExchangeRequest> | undefined, b: MessageExchangeRequest | PlainMessage<MessageExchangeRequest> | undefined): boolean;
}

/**
 * This is what will be sent from an app to the middleware
 *
 * @generated from message search.v1.AppSendRequest
 */
export declare class AppSendRequest extends Message<AppSendRequest> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * name of the recipient in the local contract
   *
   * @generated from field: string recipient = 2;
   */
  recipient: string;

  /**
   * @generated from field: search.v1.AppMessage message = 3;
   */
  message?: AppMessage;

  constructor(data?: PartialMessage<AppSendRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.AppSendRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppSendRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppSendRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppSendRequest;

  static equals(a: AppSendRequest | PlainMessage<AppSendRequest> | undefined, b: AppSendRequest | PlainMessage<AppSendRequest> | undefined): boolean;
}

/**
 * This is what will be sent from the middleware to a local app
 *
 * @generated from message search.v1.AppRecvResponse
 */
export declare class AppRecvResponse extends Message<AppRecvResponse> {
  /**
   * @generated from field: string channel_id = 1;
   */
  channelId: string;

  /**
   * name of the sender in the local contract
   *
   * @generated from field: string sender = 2;
   */
  sender: string;

  /**
   * @generated from field: search.v1.AppMessage message = 3;
   */
  message?: AppMessage;

  constructor(data?: PartialMessage<AppRecvResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.AppRecvResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppRecvResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppRecvResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppRecvResponse;

  static equals(a: AppRecvResponse | PlainMessage<AppRecvResponse> | undefined, b: AppRecvResponse | PlainMessage<AppRecvResponse> | undefined): boolean;
}

/**
 * This is the message content that is sent by the app (this is copied as-is by the middlewares)
 * TODO: we may want to use self-describing messages to have a rich type system for messages!
 * https://protobuf.dev/programming-guides/techniques/#self-description
 *
 * @generated from message search.v1.AppMessage
 */
export declare class AppMessage extends Message<AppMessage> {
  /**
   * @generated from field: string type = 1;
   */
  type: string;

  /**
   * @generated from field: bytes body = 2;
   */
  body: Uint8Array;

  constructor(data?: PartialMessage<AppMessage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "search.v1.AppMessage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AppMessage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AppMessage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AppMessage;

  static equals(a: AppMessage | PlainMessage<AppMessage> | undefined, b: AppMessage | PlainMessage<AppMessage> | undefined): boolean;
}

