---
id: alert
title: alert
---

# Synopsis

```
alert <subscommand> [option]... [args]...
```

# Description

This is a utility program for interacting with the Dexter Alerts gRPC API.  With it, one can list, add, and remove alerts.

# Subcommands

## ls

List existing alerts.

## add [file.json]

Add an alert using a JSON file.

## update [file.json]

Update an alert using a JSON file.

## get [id]

Fetch an alert by id.

## rm [id]

Remove an alert by id.

## i

List available indicators

# Examples

First, this is what a JSON file describing an alert looks like:

```json
{
  "external_id": 1,
  "exchange": "binance",
  "market": "BTC/USDT",
  "timeframe": "1m",
  "line_a": {
    "name": "Horizontal Line",
    "inputs": [10000],
    "output": "default"
  },
  "condition": 0,
  "line_b": {
    "name": "Price",
    "inputs": [],
    "output": "default"
  },
  "frequency": 0,
  "message_body": "crossed 10k",
  "webhook": {
    "method": "POST",
    "url": "http://localhost:3000/",
    "body": ""
  }
}
```

This tells dexter to send a POST request to `http://localhost:3000` if the price of BTC/USDT crosses 10000 from either direction.
Assuming the JSON above were in a file called `10k.json`, the alert would be added by doing:

```sh
bin/alert add 10k.json
```
