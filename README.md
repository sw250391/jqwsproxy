# jqwsproxy is a WebSocket proxy that supports [jq](https://stedolan.github.io/jq/) optional transformations on the incoming and outgoing WebSocket data.  It uses [websocketproxy](https://github.com/koding/websocketproxy) for the WebSocket proxy functionality and [gojq](https://github.com/itchyny/gojq) as a pure go implementation of jq

## Install

```bash
go get github.com/sw250391/jqwsproxy
```

## Examples

Listen on port `4050` and proxy connections to the server at `10.1.1.4` port `8080`.  Log the incoming and outgoing data to the terminal.

```
jqwsproxy -log=true -bind=:4050 -backend=ws://10.1.1.4:8080
```

Listen on the default port `8080` and proxy connections to the server at `10.1.1.4` port `8080`.  Log the incoming and outgoing data to the terminal (including the transformed data) and apply the specified jq query to the data being sent from the server.

```
jqwsproxy -log=true -backend=ws://10.1.1.4:8080 -inputjq=". |= ((select(.deviceType == \"Sensor\") | select(.data.responseType == \"SENSOR_CHECK_FAILURE\") | .data.success = true | .data.responseType = \"SENSOR_CHECK_SUCCESS\") // . )"
```
