# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [alerts.proto](#alerts.proto)
    - [Alert](#dexter.Alert)
    - [DeleteAlertRequest](#dexter.DeleteAlertRequest)
    - [DeleteAlertResponse](#dexter.DeleteAlertResponse)
    - [GetAlertRequest](#dexter.GetAlertRequest)
    - [Indicator](#dexter.Indicator)
    - [Line](#dexter.Line)
    - [ListAlertsRequest](#dexter.ListAlertsRequest)
    - [ListAlertsResponse](#dexter.ListAlertsResponse)
    - [ListIndicatorsRequest](#dexter.ListIndicatorsRequest)
    - [ListIndicatorsResponse](#dexter.ListIndicatorsResponse)
    - [Webhook](#dexter.Webhook)
  
    - [Condition](#dexter.Condition)
    - [Frequency](#dexter.Frequency)
  
  
    - [Alerts](#dexter.Alerts)
  

- [Scalar Value Types](#scalar-value-types)



<a name="alerts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## alerts.proto



<a name="dexter.Alert"></a>

### Alert



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| external_id | [uint64](#uint64) |  |  |
| exchange | [string](#string) |  |  |
| market | [string](#string) |  |  |
| timeframe | [string](#string) |  |  |
| line_a | [Line](#dexter.Line) |  |  |
| condition | [Condition](#dexter.Condition) |  |  |
| line_b | [Line](#dexter.Line) |  |  |
| frequency | [Frequency](#dexter.Frequency) |  |  |
| message_body | [string](#string) |  |  |
| webhook | [Webhook](#dexter.Webhook) |  |  |






<a name="dexter.DeleteAlertRequest"></a>

### DeleteAlertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alert_id | [uint64](#uint64) |  |  |






<a name="dexter.DeleteAlertResponse"></a>

### DeleteAlertResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alert_id | [uint64](#uint64) |  |  |






<a name="dexter.GetAlertRequest"></a>

### GetAlertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alert_id | [uint64](#uint64) |  |  |






<a name="dexter.Indicator"></a>

### Indicator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| inputs | [string](#string) | repeated |  |
| outputs | [string](#string) | repeated |  |






<a name="dexter.Line"></a>

### Line



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| inputs | [float](#float) | repeated |  |
| output | [string](#string) |  |  |






<a name="dexter.ListAlertsRequest"></a>

### ListAlertsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| external_id | [uint64](#uint64) |  |  |






<a name="dexter.ListAlertsResponse"></a>

### ListAlertsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alerts | [Alert](#dexter.Alert) | repeated |  |






<a name="dexter.ListIndicatorsRequest"></a>

### ListIndicatorsRequest







<a name="dexter.ListIndicatorsResponse"></a>

### ListIndicatorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| indicators | [Indicator](#dexter.Indicator) | repeated |  |






<a name="dexter.Webhook"></a>

### Webhook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [string](#string) |  |  |
| url | [string](#string) |  |  |
| body | [string](#string) |  |  |





 


<a name="dexter.Condition"></a>

### Condition


| Name | Number | Description |
| ---- | ------ | ----------- |
| Crossing | 0 |  |
| CrossingUp | 1 |  |
| CrossingDown | 2 |  |
| GreaterThan | 3 |  |
| LessThan | 4 |  |
| EnteringChannel | 5 |  |
| ExitingChannel | 6 |  |
| InsideChannel | 7 |  |
| OutsideChannel | 8 |  |
| MovingUp | 9 |  |
| MovingDown | 10 |  |
| MovingUpPercent | 11 |  |
| MovingDownPercent | 12 |  |



<a name="dexter.Frequency"></a>

### Frequency


| Name | Number | Description |
| ---- | ------ | ----------- |
| OnlyOnce | 0 |  |
| OncePerBar | 1 |  |
| OncePerBarClose | 2 |  |
| OncePerMinute | 3 |  |


 

 


<a name="dexter.Alerts"></a>

### Alerts


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAlert | [Alert](#dexter.Alert) | [Alert](#dexter.Alert) |  |
| ListAlerts | [ListAlertsRequest](#dexter.ListAlertsRequest) | [ListAlertsResponse](#dexter.ListAlertsResponse) |  |
| GetAlert | [GetAlertRequest](#dexter.GetAlertRequest) | [Alert](#dexter.Alert) |  |
| UpdateAlert | [Alert](#dexter.Alert) | [Alert](#dexter.Alert) |  |
| DeleteAlert | [DeleteAlertRequest](#dexter.DeleteAlertRequest) | [DeleteAlertResponse](#dexter.DeleteAlertResponse) |  |
| ListIndicators | [ListIndicatorsRequest](#dexter.ListIndicatorsRequest) | [ListIndicatorsResponse](#dexter.ListIndicatorsResponse) |  |

 



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

