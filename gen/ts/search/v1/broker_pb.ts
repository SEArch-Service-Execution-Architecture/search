// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file search/v1/broker.proto (package search.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { GlobalContract, LocalContract, MessageTranslations } from "./contracts_pb.js";

/**
 * Sent by client middleware to the broker to register a new channel.
 *
 * @generated from message search.v1.BrokerChannelRequest
 */
export class BrokerChannelRequest extends Message<BrokerChannelRequest> {
  /**
   * requirements contract
   *
   * @generated from field: search.v1.GlobalContract contract = 1;
   */
  contract?: GlobalContract;

  /**
   * subset of contract's participants that are already decided. This should at least
   * include the initiator's RemoteParticpant data
   *
   * @generated from field: map<string, search.v1.RemoteParticipant> preset_participants = 2;
   */
  presetParticipants: { [key: string]: RemoteParticipant } = {};

  constructor(data?: PartialMessage<BrokerChannelRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "search.v1.BrokerChannelRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contract", kind: "message", T: GlobalContract },
    { no: 2, name: "preset_participants", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: RemoteParticipant} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BrokerChannelRequest {
    return new BrokerChannelRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BrokerChannelRequest {
    return new BrokerChannelRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BrokerChannelRequest {
    return new BrokerChannelRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BrokerChannelRequest | PlainMessage<BrokerChannelRequest> | undefined, b: BrokerChannelRequest | PlainMessage<BrokerChannelRequest> | undefined): boolean {
    return proto3.util.equals(BrokerChannelRequest, a, b);
  }
}

/**
 * @generated from message search.v1.BrokerChannelResponse
 */
export class BrokerChannelResponse extends Message<BrokerChannelResponse> {
  /**
   * uuidv4
   *
   * @generated from field: string channel_id = 2;
   */
  channelId = "";

  /**
   * preset + brokered participants
   *
   * @generated from field: map<string, search.v1.RemoteParticipant> participants = 3;
   */
  participants: { [key: string]: RemoteParticipant } = {};

  /**
   * message name translations for each participant
   *
   * @generated from field: repeated search.v1.MessageTranslations messagetranslations = 4;
   */
  messagetranslations: MessageTranslations[] = [];

  constructor(data?: PartialMessage<BrokerChannelResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "search.v1.BrokerChannelResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "channel_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "participants", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: RemoteParticipant} },
    { no: 4, name: "messagetranslations", kind: "message", T: MessageTranslations, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BrokerChannelResponse {
    return new BrokerChannelResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BrokerChannelResponse {
    return new BrokerChannelResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BrokerChannelResponse {
    return new BrokerChannelResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BrokerChannelResponse | PlainMessage<BrokerChannelResponse> | undefined, b: BrokerChannelResponse | PlainMessage<BrokerChannelResponse> | undefined): boolean {
    return proto3.util.equals(BrokerChannelResponse, a, b);
  }
}

/**
 * @generated from message search.v1.RegisterProviderRequest
 */
export class RegisterProviderRequest extends Message<RegisterProviderRequest> {
  /**
   * @generated from field: search.v1.LocalContract contract = 1;
   */
  contract?: LocalContract;

  /**
   * @generated from field: string url = 2;
   */
  url = "";

  constructor(data?: PartialMessage<RegisterProviderRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "search.v1.RegisterProviderRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "contract", kind: "message", T: LocalContract },
    { no: 2, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterProviderRequest {
    return new RegisterProviderRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterProviderRequest {
    return new RegisterProviderRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterProviderRequest {
    return new RegisterProviderRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterProviderRequest | PlainMessage<RegisterProviderRequest> | undefined, b: RegisterProviderRequest | PlainMessage<RegisterProviderRequest> | undefined): boolean {
    return proto3.util.equals(RegisterProviderRequest, a, b);
  }
}

/**
 * The registry assigns the provider an ID
 *
 * @generated from message search.v1.RegisterProviderResponse
 */
export class RegisterProviderResponse extends Message<RegisterProviderResponse> {
  /**
   * @generated from field: string app_id = 1;
   */
  appId = "";

  constructor(data?: PartialMessage<RegisterProviderResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "search.v1.RegisterProviderResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "app_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterProviderResponse {
    return new RegisterProviderResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterProviderResponse {
    return new RegisterProviderResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterProviderResponse {
    return new RegisterProviderResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterProviderResponse | PlainMessage<RegisterProviderResponse> | undefined, b: RegisterProviderResponse | PlainMessage<RegisterProviderResponse> | undefined): boolean {
    return proto3.util.equals(RegisterProviderResponse, a, b);
  }
}

/**
 * @generated from message search.v1.RemoteParticipant
 */
export class RemoteParticipant extends Message<RemoteParticipant> {
  /**
   * points to the middleware for this participant
   *
   * @generated from field: string url = 1;
   */
  url = "";

  /**
   * points to the specific app that is served by the middleware
   *
   * @generated from field: string app_id = 2;
   */
  appId = "";

  constructor(data?: PartialMessage<RemoteParticipant>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "search.v1.RemoteParticipant";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "app_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoteParticipant {
    return new RemoteParticipant().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoteParticipant {
    return new RemoteParticipant().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoteParticipant {
    return new RemoteParticipant().fromJsonString(jsonString, options);
  }

  static equals(a: RemoteParticipant | PlainMessage<RemoteParticipant> | undefined, b: RemoteParticipant | PlainMessage<RemoteParticipant> | undefined): boolean {
    return proto3.util.equals(RemoteParticipant, a, b);
  }
}

