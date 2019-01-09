FROM golang:alpine as builder

ENV SRC_DIR=/go/src/github.com/AlexeyKremsa/CustomRedis
COPY . $SRC_DIR
WORKDIR $SRC_DIR
RUN cd ./cmd/restapi && go build -o customredis

FROM scratch
COPY --from=builder /go/src/github.com/AlexeyKremsa/CustomRedis/cmd/restapi/customredis customredis
CMD ["customredis"]