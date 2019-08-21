FROM node:10.16-alpine

WORKDIR /dexter-data

# sys dependency phase
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    python \
    ca-certificates

# app dependency phase
COPY package.json yarn.lock ./
RUN yarn install

# app build phase
COPY knexfile.js tsconfig.json tslint.json ./
COPY src/ src/
RUN yarn build-dist

# copy application artifacts
COPY proto/ proto/

COPY bin/ bin/

EXPOSE 50051

CMD [ "bin/dexter-data" ]
