ARG GOVERSION="1.27"
ARG USERNAME=search
FROM golang:${GOVERSION} AS dev
ARG TARGETARCH

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.60 && \
    wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-${TARGETARCH} && \
    chmod +x /usr/local/bin/grpc_health_probe

WORKDIR /src

COPY go.* .
RUN go mod download
COPY cfsm ./cfsm
COPY contract ./contract
COPY ent ./ent
COPY mocks ./mocks
COPY gen/go ./gen/go
COPY cmd ./cmd
COPY internal ./internal
COPY examples/credit-card-payments/client-go ./

RUN go build -v -o /usr/local/bin/client-go main.go
COPY examples/credit-card-payments/contract.fsa .

CMD ["/usr/local/bin/client-go"]
