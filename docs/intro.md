---
id: intro
title: Introduction
---

# What is DEXter?

DEXter is a system for running realtime market alerts.  It is implemented as two gRPC services:

* dexter-data for accessing exchange and market data via the [ccxt](https://ccxt.trade/) library.
* dexter for managing alerts

dexter-data is written in TypeScript, because ccxt is primarily a Javascript library, and dexter is
implemented in Go.  They're able to talk to each other via gRPC with dexter depending on dexter-data
for realtime price data.

# Installation

The easiest way is to let docker-compose put everything together for you.

```sh
git clone https://github.com/whiteblock/dexter.git
cd dexter
docker-compose up
```

If it's begin difficult:

```sh 
docker-compose up --force-recreate --build
```

# What do you get after doing all this?

* dexter-data providing the Data API on port 50051
* dexter providing the Alerts API on port 50052
* grpcwebproxy bridging the Data API over HTTP 1.1 on port 8081
* grpcwebproxy bridging the Alerts API over HTTP 1.1 on port 8081
* postgresql listening on port 5432

It's all described in the [docker-compose.yml](https://github.com/whiteblock/dexter/blob/master/docker-compose.yml) file.
