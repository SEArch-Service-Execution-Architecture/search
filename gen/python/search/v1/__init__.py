# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: search/v1/app_message.proto, search/v1/broker.proto, search/v1/contracts.proto, search/v1/middleware.proto
# plugin: python-betterproto
# This file has been @generated

from dataclasses import dataclass
from typing import (
    TYPE_CHECKING,
    AsyncIterable,
    AsyncIterator,
    Dict,
    Iterable,
    List,
    Optional,
    Union,
)

import betterproto
import grpclib
from betterproto.grpc.grpclib_server import ServiceBase


if TYPE_CHECKING:
    import grpclib.server
    from betterproto.grpc.grpclib_client import MetadataLike
    from grpclib.metadata import Deadline


class GlobalContractFormat(betterproto.Enum):
    UNSPECIFIED = 0
    FSA = 1
    GC = 2


class LocalContractFormat(betterproto.Enum):
    UNSPECIFIED = 0
    FSA = 1
    PYTHON_BISIMULATION_CODE = 2


class AppSendResponseResult(betterproto.Enum):
    RESULT_UNSPECIFIED = 0
    RESULT_OK = 1
    RESULT_ERR = 2


class InitChannelResponseResult(betterproto.Enum):
    RESULT_UNSPECIFIED = 0
    RESULT_ACK = 1
    RESULT_ERR = 2


class StartChannelResponseResult(betterproto.Enum):
    RESULT_UNSPECIFIED = 0
    RESULT_ACK = 1
    RESULT_ERR = 2


class CloseChannelResponseResult(betterproto.Enum):
    RESULT_UNSPECIFIED = 0
    RESULT_CLOSED = 1
    RESULT_PENDING_INBOUND = 2
    RESULT_PENDING_OUTBOUND = 3


class MessageExchangeResponseResult(betterproto.Enum):
    RESULT_UNSPECIFIED = 0
    RESULT_OK = 1
    RESULT_ERROR = 2


@dataclass(eq=False, repr=False)
class MessageExchangeRequest(betterproto.Message):
    """This is what will be exchanged between middlewares"""

    channel_id: str = betterproto.string_field(1)
    sender_id: str = betterproto.string_field(2)
    """
    This is necessary because URLs don't univocally determine apps. There can be multiple applications
     behind the same middleware (there is a 1:1 mapping between URLs and middlewares)
    """

    recipient_id: str = betterproto.string_field(3)
    content: "AppMessage" = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class AppSendRequest(betterproto.Message):
    """This is what will be sent from an app to the middleware"""

    channel_id: str = betterproto.string_field(1)
    recipient: str = betterproto.string_field(2)
    message: "AppMessage" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class AppRecvResponse(betterproto.Message):
    """This is what will be sent from the middleware to a local app"""

    channel_id: str = betterproto.string_field(1)
    sender: str = betterproto.string_field(2)
    message: "AppMessage" = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class AppMessage(betterproto.Message):
    """
    This is the message content that is sent by the app (this is copied as-is by the middlewares)
     TODO: we may want to use self-describing messages to have a rich type system for messages!
     https://protobuf.dev/programming-guides/techniques/#self-description
    """

    type: str = betterproto.string_field(1)
    body: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class GlobalContract(betterproto.Message):
    contract: bytes = betterproto.bytes_field(1)
    format: "GlobalContractFormat" = betterproto.enum_field(2)
    initiator_name: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class LocalContract(betterproto.Message):
    contract: bytes = betterproto.bytes_field(1)
    format: "LocalContractFormat" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class MessageTranslations(betterproto.Message):
    participant: str = betterproto.string_field(1)
    """
    Different contracts can be compatible while using different names for messages that are equivalent.
     This data structure contains the mapping of each message name to the name used by the other participant.
     The keys are the message names in the contract of who receives this message, and the values are the message
     names according to the other participant's contract.
    """

    translations: Dict[str, str] = betterproto.map_field(
        2, betterproto.TYPE_STRING, betterproto.TYPE_STRING
    )


@dataclass(eq=False, repr=False)
class BrokerChannelRequest(betterproto.Message):
    """Sent by client middleware to the broker to register a new channel."""

    contract: "GlobalContract" = betterproto.message_field(1)
    preset_participants: Dict[str, "RemoteParticipant"] = betterproto.map_field(
        2, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    """
    subset of contract's participants that are already decided. This should at least
     include the initiator's RemoteParticpant data
    """


@dataclass(eq=False, repr=False)
class BrokerChannelResponse(betterproto.Message):
    channel_id: str = betterproto.string_field(2)
    participants: Dict[str, "RemoteParticipant"] = betterproto.map_field(
        3, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    messagetranslations: List["MessageTranslations"] = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class RegisterProviderRequest(betterproto.Message):
    contract: "LocalContract" = betterproto.message_field(1)
    url: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class RegisterProviderResponse(betterproto.Message):
    """The registry assigns the provider an ID"""

    app_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class RemoteParticipant(betterproto.Message):
    url: str = betterproto.string_field(1)
    app_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class AppSendResponse(betterproto.Message):
    result: "AppSendResponseResult" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class AppRecvRequest(betterproto.Message):
    channel_id: str = betterproto.string_field(1)
    participant: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class RegisterChannelRequest(betterproto.Message):
    requirements_contract: "GlobalContract" = betterproto.message_field(1)
    preset_participants: Dict[str, "RemoteParticipant"] = betterproto.map_field(
        2, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class RegisterChannelResponse(betterproto.Message):
    channel_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class RegisterAppRequest(betterproto.Message):
    """
    The middleware sends this to the broker to register a provider to
       be added to the registry.
    """

    provider_contract: "LocalContract" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class RegisterAppResponse(betterproto.Message):
    """
    whenever a new channel that involves this app is started, the middleware needs to notify the local app
    """

    app_id: str = betterproto.string_field(1, group="ack_or_new")
    notification: "InitChannelNotification" = betterproto.message_field(
        2, group="ack_or_new"
    )


@dataclass(eq=False, repr=False)
class InitChannelNotification(betterproto.Message):
    """
    this is what a registered app receives whenever a new channel is initiated for that app
     the app has to communicate with the middleware using UseChannel with this new channel_id
    """

    channel_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class InitChannelRequest(betterproto.Message):
    """
    This is something that is sent by the Broker to providers to notify that a new channel is starting
    """

    channel_id: str = betterproto.string_field(1)
    app_id: str = betterproto.string_field(2)
    participants: Dict[str, "RemoteParticipant"] = betterproto.map_field(
        3, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    messagetranslations: List["MessageTranslations"] = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class InitChannelResponse(betterproto.Message):
    result: "InitChannelResponseResult" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class StartChannelRequest(betterproto.Message):
    channel_id: str = betterproto.string_field(1)
    app_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class StartChannelResponse(betterproto.Message):
    result: "StartChannelResponseResult" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class CloseChannelRequest(betterproto.Message):
    channel_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class CloseChannelResponse(betterproto.Message):
    result: "CloseChannelResponseResult" = betterproto.enum_field(1)
    error_message: str = betterproto.string_field(2)
    participants_with_pending_inbound: List[str] = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class MessageExchangeResponse(betterproto.Message):
    result: "MessageExchangeResponseResult" = betterproto.enum_field(1)
    error_message: str = betterproto.string_field(2)


class BrokerServiceStub(betterproto.ServiceStub):
    async def broker_channel(
        self,
        broker_channel_request: "BrokerChannelRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "BrokerChannelResponse":
        return await self._unary_unary(
            "/search.v1.BrokerService/BrokerChannel",
            broker_channel_request,
            BrokerChannelResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def register_provider(
        self,
        register_provider_request: "RegisterProviderRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "RegisterProviderResponse":
        return await self._unary_unary(
            "/search.v1.BrokerService/RegisterProvider",
            register_provider_request,
            RegisterProviderResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )


class PrivateMiddlewareServiceStub(betterproto.ServiceStub):
    async def register_channel(
        self,
        register_channel_request: "RegisterChannelRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "RegisterChannelResponse":
        return await self._unary_unary(
            "/search.v1.PrivateMiddlewareService/RegisterChannel",
            register_channel_request,
            RegisterChannelResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def register_app(
        self,
        register_app_request: "RegisterAppRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> AsyncIterator[RegisterAppResponse]:
        async for response in self._unary_stream(
            "/search.v1.PrivateMiddlewareService/RegisterApp",
            register_app_request,
            RegisterAppResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        ):
            yield response

    async def close_channel(
        self,
        close_channel_request: "CloseChannelRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "CloseChannelResponse":
        return await self._unary_unary(
            "/search.v1.PrivateMiddlewareService/CloseChannel",
            close_channel_request,
            CloseChannelResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def app_send(
        self,
        app_send_request: "AppSendRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "AppSendResponse":
        return await self._unary_unary(
            "/search.v1.PrivateMiddlewareService/AppSend",
            app_send_request,
            AppSendResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def app_recv(
        self,
        app_recv_request: "AppRecvRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "AppRecvResponse":
        return await self._unary_unary(
            "/search.v1.PrivateMiddlewareService/AppRecv",
            app_recv_request,
            AppRecvResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )


class PublicMiddlewareServiceStub(betterproto.ServiceStub):
    async def init_channel(
        self,
        init_channel_request: "InitChannelRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "InitChannelResponse":
        return await self._unary_unary(
            "/search.v1.PublicMiddlewareService/InitChannel",
            init_channel_request,
            InitChannelResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def start_channel(
        self,
        start_channel_request: "StartChannelRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "StartChannelResponse":
        return await self._unary_unary(
            "/search.v1.PublicMiddlewareService/StartChannel",
            start_channel_request,
            StartChannelResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def message_exchange(
        self,
        message_exchange_request_iterator: Union[
            AsyncIterable[MessageExchangeRequest], Iterable[MessageExchangeRequest]
        ],
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "MessageExchangeResponse":
        return await self._stream_unary(
            "/search.v1.PublicMiddlewareService/MessageExchange",
            message_exchange_request_iterator,
            MessageExchangeRequest,
            MessageExchangeResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )


class BrokerServiceBase(ServiceBase):

    async def broker_channel(
        self, broker_channel_request: "BrokerChannelRequest"
    ) -> "BrokerChannelResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def register_provider(
        self, register_provider_request: "RegisterProviderRequest"
    ) -> "RegisterProviderResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_broker_channel(
        self,
        stream: "grpclib.server.Stream[BrokerChannelRequest, BrokerChannelResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.broker_channel(request)
        await stream.send_message(response)

    async def __rpc_register_provider(
        self,
        stream: "grpclib.server.Stream[RegisterProviderRequest, RegisterProviderResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.register_provider(request)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/search.v1.BrokerService/BrokerChannel": grpclib.const.Handler(
                self.__rpc_broker_channel,
                grpclib.const.Cardinality.UNARY_UNARY,
                BrokerChannelRequest,
                BrokerChannelResponse,
            ),
            "/search.v1.BrokerService/RegisterProvider": grpclib.const.Handler(
                self.__rpc_register_provider,
                grpclib.const.Cardinality.UNARY_UNARY,
                RegisterProviderRequest,
                RegisterProviderResponse,
            ),
        }


class PrivateMiddlewareServiceBase(ServiceBase):

    async def register_channel(
        self, register_channel_request: "RegisterChannelRequest"
    ) -> "RegisterChannelResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def register_app(
        self, register_app_request: "RegisterAppRequest"
    ) -> AsyncIterator[RegisterAppResponse]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)
        yield RegisterAppResponse()

    async def close_channel(
        self, close_channel_request: "CloseChannelRequest"
    ) -> "CloseChannelResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def app_send(self, app_send_request: "AppSendRequest") -> "AppSendResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def app_recv(self, app_recv_request: "AppRecvRequest") -> "AppRecvResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_register_channel(
        self,
        stream: "grpclib.server.Stream[RegisterChannelRequest, RegisterChannelResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.register_channel(request)
        await stream.send_message(response)

    async def __rpc_register_app(
        self, stream: "grpclib.server.Stream[RegisterAppRequest, RegisterAppResponse]"
    ) -> None:
        request = await stream.recv_message()
        await self._call_rpc_handler_server_stream(
            self.register_app,
            stream,
            request,
        )

    async def __rpc_close_channel(
        self, stream: "grpclib.server.Stream[CloseChannelRequest, CloseChannelResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.close_channel(request)
        await stream.send_message(response)

    async def __rpc_app_send(
        self, stream: "grpclib.server.Stream[AppSendRequest, AppSendResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.app_send(request)
        await stream.send_message(response)

    async def __rpc_app_recv(
        self, stream: "grpclib.server.Stream[AppRecvRequest, AppRecvResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.app_recv(request)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/search.v1.PrivateMiddlewareService/RegisterChannel": grpclib.const.Handler(
                self.__rpc_register_channel,
                grpclib.const.Cardinality.UNARY_UNARY,
                RegisterChannelRequest,
                RegisterChannelResponse,
            ),
            "/search.v1.PrivateMiddlewareService/RegisterApp": grpclib.const.Handler(
                self.__rpc_register_app,
                grpclib.const.Cardinality.UNARY_STREAM,
                RegisterAppRequest,
                RegisterAppResponse,
            ),
            "/search.v1.PrivateMiddlewareService/CloseChannel": grpclib.const.Handler(
                self.__rpc_close_channel,
                grpclib.const.Cardinality.UNARY_UNARY,
                CloseChannelRequest,
                CloseChannelResponse,
            ),
            "/search.v1.PrivateMiddlewareService/AppSend": grpclib.const.Handler(
                self.__rpc_app_send,
                grpclib.const.Cardinality.UNARY_UNARY,
                AppSendRequest,
                AppSendResponse,
            ),
            "/search.v1.PrivateMiddlewareService/AppRecv": grpclib.const.Handler(
                self.__rpc_app_recv,
                grpclib.const.Cardinality.UNARY_UNARY,
                AppRecvRequest,
                AppRecvResponse,
            ),
        }


class PublicMiddlewareServiceBase(ServiceBase):

    async def init_channel(
        self, init_channel_request: "InitChannelRequest"
    ) -> "InitChannelResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def start_channel(
        self, start_channel_request: "StartChannelRequest"
    ) -> "StartChannelResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def message_exchange(
        self, message_exchange_request_iterator: AsyncIterator[MessageExchangeRequest]
    ) -> "MessageExchangeResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_init_channel(
        self, stream: "grpclib.server.Stream[InitChannelRequest, InitChannelResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.init_channel(request)
        await stream.send_message(response)

    async def __rpc_start_channel(
        self, stream: "grpclib.server.Stream[StartChannelRequest, StartChannelResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.start_channel(request)
        await stream.send_message(response)

    async def __rpc_message_exchange(
        self,
        stream: "grpclib.server.Stream[MessageExchangeRequest, MessageExchangeResponse]",
    ) -> None:
        request = stream.__aiter__()
        response = await self.message_exchange(request)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/search.v1.PublicMiddlewareService/InitChannel": grpclib.const.Handler(
                self.__rpc_init_channel,
                grpclib.const.Cardinality.UNARY_UNARY,
                InitChannelRequest,
                InitChannelResponse,
            ),
            "/search.v1.PublicMiddlewareService/StartChannel": grpclib.const.Handler(
                self.__rpc_start_channel,
                grpclib.const.Cardinality.UNARY_UNARY,
                StartChannelRequest,
                StartChannelResponse,
            ),
            "/search.v1.PublicMiddlewareService/MessageExchange": grpclib.const.Handler(
                self.__rpc_message_exchange,
                grpclib.const.Cardinality.STREAM_UNARY,
                MessageExchangeRequest,
                MessageExchangeResponse,
            ),
        }
