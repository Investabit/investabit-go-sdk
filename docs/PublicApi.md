# \PublicApi

All URIs are relative to *https://api.investabit.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PublicCurrentSymbolGet**](PublicApi.md#V1PublicCurrentSymbolGet) | **Get** /v1/public/current/{symbol} | Current
[**V1PublicPriceHistorySymbolGet**](PublicApi.md#V1PublicPriceHistorySymbolGet) | **Get** /v1/public/price-history/{symbol} | Price History
[**V1PublicSymbolsGet**](PublicApi.md#V1PublicSymbolsGet) | **Get** /v1/public/symbols | Symbols
[**V1PublicTrendSymbolGet**](PublicApi.md#V1PublicTrendSymbolGet) | **Get** /v1/public/trend/{symbol} | Trend


# **V1PublicCurrentSymbolGet**
> PublicCurrentResponse V1PublicCurrentSymbolGet(ctx, symbol)
Current



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol, default is btc. | 

### Return type

[**PublicCurrentResponse**](Public Current Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PublicPriceHistorySymbolGet**
> PublicPriceResponse V1PublicPriceHistorySymbolGet(ctx, symbol)
Price History



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol, default is btc. | 

### Return type

[**PublicPriceResponse**](Public Price Response.md)

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
  **symbol** | **string**| The cryptocurrency symbol, default is btc. | 

### Return type

[**PublicTrendResponse**](Public Trend Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

