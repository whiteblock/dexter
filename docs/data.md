---
id: data
title: Data gRPC API
---
# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [data.proto](#data.proto)
    - [Candle](#dexter.Candle)
    - [CandlesRequest](#dexter.CandlesRequest)
    - [CandlesResponse](#dexter.CandlesResponse)
    - [ExchangesRequest](#dexter.ExchangesRequest)
    - [ExchangesResponse](#dexter.ExchangesResponse)
    - [MarketsRequest](#dexter.MarketsRequest)
    - [MarketsResponse](#dexter.MarketsResponse)
    - [TestRequest](#dexter.TestRequest)
    - [TestResponse](#dexter.TestResponse)
  
  
  
    - [Data](#dexter.Data)
  

- [Scalar Value Types](#scalar-value-types)



<a name="data.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## data.proto



<a name="dexter.Candle"></a>

### Candle



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [uint64](#uint64) |  |  |
| o | [double](#double) |  |  |
| h | [double](#double) |  |  |
| l | [double](#double) |  |  |
| c | [double](#double) |  |  |
| v | [double](#double) |  |  |






<a name="dexter.CandlesRequest"></a>

### CandlesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange | [string](#string) |  |  |
| market | [string](#string) |  |  |
| timeframe | [string](#string) |  |  |
| since | [string](#string) |  |  |
| limit | [uint64](#uint64) |  |  |






<a name="dexter.CandlesResponse"></a>

### CandlesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candles | [Candle](#dexter.Candle) | repeated |  |






<a name="dexter.ExchangesRequest"></a>

### ExchangesRequest







<a name="dexter.ExchangesResponse"></a>

### ExchangesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchanges | [string](#string) | repeated |  |






<a name="dexter.MarketsRequest"></a>

### MarketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange | [string](#string) |  |  |






<a name="dexter.MarketsResponse"></a>

### MarketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| markets | [string](#string) | repeated |  |






<a name="dexter.TestRequest"></a>

### TestRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| a | [string](#string) |  |  |






<a name="dexter.TestResponse"></a>

### TestResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| a | [int64](#int64) |  |  |
| b | [int64](#int64) |  |  |





 

 

 


<a name="dexter.Data"></a>

### Data


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SupportedExchanges | [ExchangesRequest](#dexter.ExchangesRequest) | [ExchangesResponse](#dexter.ExchangesResponse) |  |
| SupportedMarkets | [MarketsRequest](#dexter.MarketsRequest) | [MarketsResponse](#dexter.MarketsResponse) |  |
| GetCandles | [CandlesRequest](#dexter.CandlesRequest) | [CandlesResponse](#dexter.CandlesResponse) |  |
| StreamCandles | [CandlesRequest](#dexter.CandlesRequest) | [Candle](#dexter.Candle) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

