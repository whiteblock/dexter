FROM golang:1.12.6-alpine

WORKDIR /dexter

# sys dependency phase
RUN apk add --no-cache git make

# app dependency phase
COPY go.mod go.sum ./
RUN go get

# app build phase
COPY *.go Makefile ./
COPY api/ api/
COPY cmd/ cmd/
COPY demo/ demo/

RUN make

EXPOSE 50052
CMD [ "./dexter" ]


