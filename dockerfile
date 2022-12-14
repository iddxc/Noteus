FROM golang:alpine as builder

WORKDIR /go/src/Noteus
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

WORKDIR /go/src/Noteus

COPY --from=builder /go/src/Noteus ./
COPY --from=builder /go/src/Noteus/config.yaml ./

ENTRYPOINT ./server
