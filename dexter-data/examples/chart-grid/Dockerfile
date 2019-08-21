FROM node:10.16-alpine

WORKDIR /chart-grid

# sys dependency phase
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    python
RUN apk add --no-cache ca-certificates libressl && update-ca-certificates

# app dependency phase
COPY package.json yarn.lock next.config.js ./
COPY components/ components/
COPY lib/ lib/
COPY pages/ pages/
COPY static/ static/
RUN yarn install

# app build phase
RUN yarn build

EXPOSE 3000

CMD [ "yarn", "start" ]
