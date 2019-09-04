---
id: dexter
title: dexter
---

## Synopsis

```sh
Usage: dexter [OPTION]...
```

## Description

This is responsible for running the Dexter Alerts gRPC service which allows one to set price alerts on cryptocurrency markets.

## Options

### -h, --help

Display help message

### -v, --verbose

Be verbose

### -c, --client [address]

Bind address of the dexter-data service to consume.  Defaults to ("0.0.0.0:50051")

### -l, --listen [address]

Bind address for this dexter service.  Defaults to ("0.0.0.0:50052")

## Environment Variables

### PG_URL

This should be a [PostgreSQL connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING).

https://stackoverflow.com/questions/3582552/postgresql-connection-url

## Examples

**Start dexter**

```
./dexter
```
