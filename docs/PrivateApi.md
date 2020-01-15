# \PrivateApi

All URIs are relative to *https://api.cryptoweather.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PrivateAccuracySymbolIntervalPeriodGet**](PrivateApi.md#V1PrivateAccuracySymbolIntervalPeriodGet) | **Get** /v1/private/accuracy/{symbol}/{interval}/{period} | Accuracy
[**V1PrivateForecastSymbolIntervalGet**](PrivateApi.md#V1PrivateForecastSymbolIntervalGet) | **Get** /v1/private/forecast/{symbol}/{interval} | Forecast
[**V1PrivateTrendSymbolGet**](PrivateApi.md#V1PrivateTrendSymbolGet) | **Get** /v1/private/trend/{symbol} | Trend
[**V1PrivateTrendTabularGet**](PrivateApi.md#V1PrivateTrendTabularGet) | **Get** /v1/private/trend-tabular | Trend Tabular


# **V1PrivateAccuracySymbolIntervalPeriodGet**
> PrivateAccuracyResponse V1PrivateAccuracySymbolIntervalPeriodGet(ctx, symbol, interval, period, optional)
Accuracy

The accuracy response contains the following attributes.  + `rmse` Root Mean Square Error  + `mae` Mean Absolute error  + `r2` R Squared  + `ci` Confidence Interval lower and upper error bounds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol. | 
  **interval** | **string**| The forecast interval, 1h or 1d. | 
  **period** | **string**| The period for computing the accuracy, such as the past 7 days. | 
 **optional** | ***V1PrivateAccuracySymbolIntervalPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1PrivateAccuracySymbolIntervalPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cookie** | **optional.String**| e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89 | 
 **xCsrf** | **optional.String**| e.g. b1820141-1bad-4a9c-93c0-52022817ce89 | 

### Return type

[**PrivateAccuracyResponse**](Private Accuracy Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PrivateForecastSymbolIntervalGet**
> PrivateForecastResponse V1PrivateForecastSymbolIntervalGet(ctx, symbol, interval, optional)
Forecast

The forecast response contains a sequence of forecasts at the specified intervals, with the following attributes.  + `time_start` start time of the period the forecast is applicable for  + `time_end` end time of the period the forecast is applicable for  + `low` forecasted high during the period  + `high` forecasted low during the period  + `weighted_price` forecasted weighted price during the period  + `change_pct` percent change in price for forecasted weighted_price relative to current price  + `change_usd` dollar change in price for forecasted weighted_price relative to current price

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol. | 
  **interval** | **string**| The forecast interval, 1h or 1d. | 
 **optional** | ***V1PrivateForecastSymbolIntervalGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1PrivateForecastSymbolIntervalGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cookie** | **optional.String**| e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89 | 
 **xCsrf** | **optional.String**| e.g. b1820141-1bad-4a9c-93c0-52022817ce89 | 

### Return type

[**PrivateForecastResponse**](Private Forecast Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PrivateTrendSymbolGet**
> PublicTrendResponse V1PrivateTrendSymbolGet(ctx, symbol, optional)
Trend

The trend response contains a collection of forecasts for different intervals with the following attributes.  + `time_start` start time of the period the forecast is applicable for  + `time_end` end time of the period the forecast is applicable for  + `interval` interval in hours that the forecast is applicable for  + `weighted_price` forecasted weighted price during the period  + `change_pct` percent change in price for forecasted weighted_price relative to current price  + `change_usd` dollar change in price for forecasted weighted_price relative to current price

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| The cryptocurrency symbol. | 
 **optional** | ***V1PrivateTrendSymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1PrivateTrendSymbolGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89 | 
 **xCsrf** | **optional.String**| e.g. b1820141-1bad-4a9c-93c0-52022817ce89 | 

### Return type

[**PublicTrendResponse**](Public Trend Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PrivateTrendTabularGet**
> PrivateTrendTabularResponse V1PrivateTrendTabularGet(ctx, optional)
Trend Tabular

The trend tabular response contains a collection of forecasts for all supported cryptocurrencies at different intervals with the following attributes.  + `time_start` start time of the period the forecast is applicable for  + `time_end` end time of the period the forecast is applicable for  + `interval` interval in hours that the forecast is applicable for  + `weighted_price` forecasted weighted price during the period  + `change_pct` percent change in price for forecasted weighted_price relative to current price  + `change_usd` dollar change in price for forecasted weighted_price relative to current price

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1PrivateTrendTabularGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1PrivateTrendTabularGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89 | 
 **xCsrf** | **optional.String**| e.g. b1820141-1bad-4a9c-93c0-52022817ce89 | 

### Return type

[**PrivateTrendTabularResponse**](Private Trend Tabular Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

