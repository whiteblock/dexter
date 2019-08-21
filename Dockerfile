FROM golang:1.12.6-alpine

WORKDIR /dexter

# sys dependency phase
RUN apk add --no-cache git make

# app dependency phase
COPY go.mod go.sum ./
RUN go get
RUN curl -s https://github.com/grpc/grpc-web/releases/download/1.0.6/protoc-gen-grpc-web-1.0.6-linux-x86_64 -o /usr/local/bin/protoc-gen-grpc-web
RUN chmod a+x /usr/local/bin/protoc-gen-grpc-web

# app build phase
COPY *.go Makefile ./
COPY api/ api/
COPY cmd/ cmd/

RUN make

EXPOSE 50052
CMD [ "./dexter" ]


