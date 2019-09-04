---
id: dexterdata
title: dexter-data
---

# Synopsis

```
dexter-data [OPTION]...
```

# Description

This starts the Dexter Data gRPC service which serves price data to clients. It
can be thought as a gRPC wrapper around the subset of ccxt that provides price
data over time.

# Options

### -h, --help

Display help message

### -v, --verbose

Be verbose

### -V, --version

Display the version number

### -b, --bind [address]

Bind address for this dexter service.  Defaults to ("0.0.0.0:50051"

# Examples

**Start dexter-data**

```
cd dexter-data
bin/dexter-data
```
