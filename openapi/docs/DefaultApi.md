# \DefaultApi

All URIs are relative to *http://myazure.url.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataGet**](DefaultApi.md#DataGet) | **Get** /data | Returns a string with an explanation on how to use this url.
[**DataPost**](DefaultApi.md#DataPost) | **Post** /data | Compute temperature and humidity stuff.



## DataGet

> string DataGet(ctx).Execute()

Returns a string with an explanation on how to use this url.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DataGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataGet`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DataGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPost

> Sensorefficiency DataPost(ctx).Sensordata(sensordata).Execute()

Compute temperature and humidity stuff.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sensordata := *openapiclient.NewSensordata(float32(123), float32(123)) // Sensordata | retrieve (temperature/humidity) value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DataPost(context.Background()).Sensordata(sensordata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataPost`: Sensorefficiency
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sensordata** | [**Sensordata**](Sensordata.md) | retrieve (temperature/humidity) value | 

### Return type

[**Sensorefficiency**](Sensorefficiency.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

