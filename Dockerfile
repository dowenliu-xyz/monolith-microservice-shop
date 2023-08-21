ARG GOPROXY
ARG BIN

FROM golangci/golangci-lint:v1.54.1 as golangci-lint

WORKDIR /go/src
COPY . .

RUN golangci-lint run -v

FROM golang:1.21.0 as builder
ARG GOPROXY
ARG BIN

WORKDIR /go/src/github.com/ThreeDotsLabs/monolith-microservice-shop
COPY . .

ENV GOPROXY=${GOPROXY}

RUN go mod download
RUN go install github.com/roblaszczak/go-cleanarch@latest && go-cleanarch
RUN CGO_ENABLED=0 go build -o /go/bin/${BIN} ./cmd/${BIN}

FROM ubuntu:latest
ARG BIN

WORKDIR /app
COPY --from=builder /go/bin/${BIN} /app/run

CMD ["/app/run"]
