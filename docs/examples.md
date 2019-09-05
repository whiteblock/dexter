---
id: examples
title: Examples
---

# Getting Everything Running

To make things easier for everyone, a `docker-compose.yml` has been provided that gets everything running
in one easy step.

```sh
git clone https://github.com/whiteblock/dexter.git
cd dexter
docker-compose up 
```

This gives you the following:

* PostgreSQL running on port 5432
* dexter-data running on port 50051
* grpcwebproxy running on port 8081 as a proxy for dexter-data
* dexter running on port 50052
* grpcwebproxy running on port 8082 as a proxy for dexter

The first time will take a while to download all the dependencies, but subsequent startup times should be relatively quick.

## Trying the Streaming Demo

```sh
cd demo
yarn
yarn start
```

Then visit http://localhost:3000 to see the 1 minute BTC/USDT chart from Binance.

![streaming demo](/dexter/img/streaming-demo.png)

## Setting Alerts from the Command Line

Dexter uses webhooks to send alert notifications, so for this demo, we can use the [webhook](https://github.com/adnanh/webhook** program
to provide a webhook server to receive our notifications and perform an action.

### Setup

#### Install Prequisites

```sh
sudo apt-get install wget webhook alsa-utils
```

#### Download a .wav file

```sh
cd /tmp
wget http://www.wavsource.com/snds_2018-06-03_5106726768923853/sfx/boxing_bell_multiple.wav
```

#### Create a shell script that will play the wav file

```sh
#!/bin/sh
/usr/bin/aplay /tmp/boxing_bell_multiple.wav
```

#### Create a /tmp/hooks.json

```json
[
  {
    "id": "dexter-notification",
    "execute-command": "/tmp/play.sh",
    "command-working-directory": "/tmp"
  }
]
```

#### Run webhook


```sh
webhook -hooks /tmp/hooks.json -verbose -port 9001
```

In another terminal, you can test the webhook by running:

```sh
curl -X POST http://localhost:9001/hooks/dexter-notification
```

### Setting up an Alert

#### Horizontal Line is the simplest alert

Suppose you wanted to know when BTC/USDT crossed $10000 from either direction.
The following is an alert on a **Horizontal Line**.

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
    "url": "http://localhost:9001/",
    "body": ""
  }
}
```

Put this JSON in a file called `/tmp/btc10k.json` and add the alert by doing:

```sh
cd dexter-data
bin/alert add /tmp/btc10k.json
```

#### Moving Average Crosses

Alts are tanking hard in the bear market, but that doesn't mean they can't have significant dead cat bounces
along the way.  One way to detect these situations is to use Moving Average crosses.  In this exapmle, we'll
look for the 4 hour 50-period **Moving Average** to cross up against the 4 hour 200-period **Moving Average**.
This is commonly called a 4 hour golden cross.

```json
{
  "external_id": 1,
  "exchange": "binance",
  "market": "ADA/USDT",
  "timeframe": "4h",
  "line_a": {
    "name": "Moving Average",
    "inputs": [50],
    "output": "default"
  },
  "condition": 1,
  "line_b": {
    "name": "Moving Average",
    "inputs": [200],
    "output": "default"
  },
  "frequency": 0,
  "message_body": "golden cross",
  "webhook": {
    "method": "POST",
    "url": "http://localhost:9001/hooks/dexter-notification",
    "body": ""
  }
}
```

In this json `line_a` describes a 50 MA and `line_b` describes a 200 MA.  `condition: 1` means it's looking
for `line_a` to cross up `line_b` and it comes from the `Conditions` enum defined in alerts.proto.

Put this in a file called `/tmp/golden-cross.json` and add the alert by doing:

```sh 
bin/alert add /tmp/golden-cross.json
```

