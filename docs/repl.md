---
id: repl
title: repl
---

# Synopsis

Usage:

```sh
cd dexter-data
bin/repl
```

# Description

This is a helper program for starting up a node.js REPL with gRPC client libraries preloaded and ready to interact with dexter.

# Global Objects

## dd

This namespace contains classes and functions that wrap the subset of the ccxt
library that deals with prices and market information. The dexter-data gRPC
service can be considered a wrapper around this code.

## service

This namespace provides a few functions for creating gRPC clients and servers
for the Dexter Data gRPC service.

## alerts

This namespace provides functions for creating a gRPC client for the Dexter
Alerts gRPC service.

## cl

`cl` is an alias for `console.log** which is handy for debugging in the repl.

# Examples

### Start the repl

```sh
cd dexter-data
bin/repl
```

### A typical repl session

```js
// instantiate a dexter-data gRPC client
c = service.getClient('0.0.0.0:50051')

// list the exchanges we can access
c.supportedExchanges({}, cl)

// list the markets we can access from binance
c.supportedMarkets({ exchange: 'binance' }, cl)
```
