FROM golang:1.12.6-alpine

WORKDIR /dexter

# sys dependency phase
RUN apk add --no-cache git

# app dependency phase
COPY go.mod go.sum ./
RUN go get

# app build phase
#COPY knexfile.js tsconfig.json tslint.json ./
#COPY src/ src/
#RUN yarn build-dist
#
## copy application artifacts
#COPY proto/ proto/
#
#COPY bin/ bin/
#
#EXPOSE 50052
#
#CMD [ "bin/dexter-data" ]
