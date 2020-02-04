# \PublicApi

All URIs are relative to *https://api.cryptoweather.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PublicPriceChangeSymbolGet**](PublicApi.md#V1PublicPriceChangeSymbolGet) | **Get** /v1/public/price-change/{symbol} | Price Change
[**V1PublicPriceCurrentSymbolGet**](PublicApi.md#V1PublicPriceCurrentSymbolGet) | **Get** /v1/public/price-current/{symbol} | Price Current
[**V1PublicPriceHistorySymbolPeriodIntervalGet**](PublicApi.md#V1PublicPriceHistorySymbolPeriodIntervalGet) | **Get** /v1/public/price-history/{symbol}/{period}/{interval} | Price History
[**V1PublicSummaryGet**](PublicApi.md#V1PublicSummaryGet) | **Get** /v1/public/summary | Summary
[**V1PublicSymbolsGet**](PublicApi.md#V1PublicSymbolsGet) | **Get** /v1/public/symbols | Symbols
[**V1PublicTrendSymbolGet**](PublicApi.md#V1PublicTrendSymbolGet) | **Get** /v1/public/trend/{symbol} | Trend


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

# **V1PublicSummaryGet**
> PublicSummaryResponse V1PublicSummaryGet(ctx, )
Summary



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSummaryResponse**](Public Summary Response.md)

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

The trend response contains a collection of forecasts for different intervals with the following attributes.  + `time_start` start time of the period the forecast is applicable for  + `time_end` end time of the period the forecast is applicable for  + `interval` interval in hours that the forecast is applicable for  + `weighted_price` forecasted weighted price during the period  + `change_pct` percent change in price for forecasted weighted_price relative to current price  + `change_usd` dollar change in price for forecasted weighted_price relative to current price

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

