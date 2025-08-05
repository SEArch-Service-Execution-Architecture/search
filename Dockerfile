ARG GOVERSION="1.24"
ARG USERNAME=search
FROM golang:${GOVERSION} AS dev

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.39 && \
    wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /usr/local/bin/grpc_health_probe

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /src
# ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
RUN go mod download

ENV GOCACHE=/home/$USERNAME/.cache/go-build
COPY cfsm ./cfsm
COPY contract ./contract
COPY ent ./ent
COPY mocks ./mocks
COPY gen/go ./gen/go
COPY cmd ./cmd
COPY internal ./internal

RUN --mount=type=cache,target=/home/$USERNAME/.cache/go-build \
    go build -v -o /usr/local/bin/broker cmd/broker/broker.go
RUN --mount=type=cache,target=/home/$USERNAME/.cache/go-build \
    go build -v -o /usr/local/bin/middleware cmd/middleware/middleware.go

# FROM scratch AS prod
# ENV PATH=/
# COPY --from=dev /usr/local/bin/grpc_health_probe /
# COPY --from=dev /usr/local/bin/middleware /
# COPY --from=dev /usr/local/bin/broker /


# Use Python 3.13 as base image
FROM python:3.12-slim AS with-python-vfsm-bisimulation
ENV PYTHONUNBUFFERED=1
# https://docs.astral.sh/uv/guides/integration/docker/#installing-uv
COPY --from=ghcr.io/astral-sh/uv:0.8.4 /uv /uvx /bin/
ENV UV_COMPILE_BYTECODE=1

# Set working directory to /app
WORKDIR /app

RUN uv pip install --system z3-solver cfsm-bisimulation

COPY --from=dev /usr/local/bin/grpc_health_probe /usr/local/bin/
COPY --from=dev /usr/local/bin/middleware /usr/local/bin/
COPY --from=dev /usr/local/bin/broker /usr/local/bin/

# FROM with-python-vfsm-bisimulation AS debug

# COPY --from=dev /go/bin/dlv /usr/local/bin/
# # Copy source code for debugging
# COPY --from=dev /src /src