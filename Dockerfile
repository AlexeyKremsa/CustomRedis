FROM golang:1.8

ENV SRC_DIR=/go/src/github.com/AlexeyKremsa/CustomRedis

WORKDIR $SRC_DIR
COPY . $SRC_DIR
RUN cd $SRC_DIR; go build
CMD ["./CustomRedis"]