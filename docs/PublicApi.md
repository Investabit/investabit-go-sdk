# \PublicApi

All URIs are relative to *https://api.investabit.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PublicPriceChangeSymbolGet**](PublicApi.md#V1PublicPriceChangeSymbolGet) | **Get** /v1/public/price-change/{symbol} | Price Change
[**V1PublicPriceCurrentSymbolGet**](PublicApi.md#V1PublicPriceCurrentSymbolGet) | **Get** /v1/public/price-current/{symbol} | Price Current
[**V1PublicPriceHistorySymbolPeriodIntervalGet**](PublicApi.md#V1PublicPriceHistorySymbolPeriodIntervalGet) | **Get** /v1/public/price-history/{symbol}/{period}/{interval} | Price History
[**V1PublicSymbolsGet**](PublicApi.md#V1PublicSymbolsGet) | **Get** /v1/public/symbols | Symbols
[**V1PublicTrendSymbolGet**](PublicApi.md#V1PublicTrendSymbolGet) | **Get** /v1/public/trend/{symbol} | Trend
[**V1PublicWaitingListPost**](PublicApi.md#V1PublicWaitingListPost) | **Post** /v1/public/waiting-list | Waiting List


# **V1PublicPriceChangeSymbolGet**
> PublicPriceChangeResponse V1PublicPriceChangeSymbolGet(ctx, symbol)
Price Change



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol. | 

### Return type

[**PublicPriceChangeResponse**](Public Price Change Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicPriceCurrentSymbolGet**
> PublicPriceCurrentResponse V1PublicPriceCurrentSymbolGet(ctx, symbol)
Price Current



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol, provide &#x60;all&#x60; to get every symbol. | 

### Return type

[**PublicPriceCurrentResponse**](Public Price Current Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicPriceHistorySymbolPeriodIntervalGet**
> PublicPriceHistoryResponse V1PublicPriceHistorySymbolPeriodIntervalGet(ctx, symbol, period, interval)
Price History



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol, provide &#x60;all&#x60; to get every symbol. | 
  **period** | **string**| The period to get data for, such as past 30 days. | 
  **interval** | **string**| The bar interval, such as 1 day. | 

### Return type

[**PublicPriceHistoryResponse**](Public Price History Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicSymbolsGet**
> PublicSymbolsResponse V1PublicSymbolsGet(ctx, )
Symbols



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSymbolsResponse**](Public Symbols Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicTrendSymbolGet**
> PublicTrendResponse V1PublicTrendSymbolGet(ctx, symbol)
Trend



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol. | 

### Return type

[**PublicTrendResponse**](Public Trend Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicWaitingListPost**
> DefaultResponse V1PublicWaitingListPost(ctx, optional)
Waiting List

Subscribe a user to the waiting list, `name` is not required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1PublicWaitingListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1PublicWaitingListPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PublicWaitingListRequest**](PublicWaitingListRequest.md)|  | 

### Return type

[**DefaultResponse**](Default Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

