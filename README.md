# Go API client for swagger

The Investabit API allows for access to all of the public facing services provided, including market data and forecasts.  ## General Overview  1. All API methods will be built to adhere to RESTful best practices as closely as possible. As such, all API calls will be made via the standard HTTP protocol using the GET/POST/PUT/DELETE request types.  2. Every request returns the status as a JSON response with the following   - success, true if it was successful   - code, the http status code (also in the response header)          200 if response is successful          400 if bad request          401 if authorization JWT is wrong or limit exceeded          404 wrong route          500 for any internal errors  - status, the status of the request, default **success**  - errors, an array of any relevant error details  3. For any requests that are not successful an error message is specified and returned as an array for the **errors** key in the JSON response.  4. All authentication uses JSON Web Tokens (JWT), which is set as the **Authorization** entry in the header, see the following for more details.     * http://jwt.io     * https://scotch.io/tutorials/the-anatomy-of-a-json-web-token

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.investabit.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PublicApi* | [**V1PublicCurrentSymbolGet**](docs/PublicApi.md#v1publiccurrentsymbolget) | **Get** /v1/public/current/{symbol} | Current
*PublicApi* | [**V1PublicPriceChangeSymbolGet**](docs/PublicApi.md#v1publicpricechangesymbolget) | **Get** /v1/public/price-change/{symbol} | Price Change
*PublicApi* | [**V1PublicPriceHistorySymbolPeriodIntervalGet**](docs/PublicApi.md#v1publicpricehistorysymbolperiodintervalget) | **Get** /v1/public/price-history/{symbol}/{period}/{interval} | Price History
*PublicApi* | [**V1PublicSymbolsGet**](docs/PublicApi.md#v1publicsymbolsget) | **Get** /v1/public/symbols | Symbols
*PublicApi* | [**V1PublicTrendSymbolGet**](docs/PublicApi.md#v1publictrendsymbolget) | **Get** /v1/public/trend/{symbol} | Trend


## Documentation For Models

 - [CurrentRoute](docs/CurrentRoute.md)
 - [DefaultResponse](docs/DefaultResponse.md)
 - [PriceChangeRoute](docs/PriceChangeRoute.md)
 - [PriceHistoryRoute](docs/PriceHistoryRoute.md)
 - [PublicCurrentResponse](docs/PublicCurrentResponse.md)
 - [PublicCurrentResponseData](docs/PublicCurrentResponseData.md)
 - [PublicCurrentResponseDataCurrent](docs/PublicCurrentResponseDataCurrent.md)
 - [PublicPriceChangeResponse](docs/PublicPriceChangeResponse.md)
 - [PublicPriceChangeResponseData](docs/PublicPriceChangeResponseData.md)
 - [PublicPriceChangeResponseDataPriceChange](docs/PublicPriceChangeResponseDataPriceChange.md)
 - [PublicPriceHistoryResponse](docs/PublicPriceHistoryResponse.md)
 - [PublicPriceHistoryResponseData](docs/PublicPriceHistoryResponseData.md)
 - [PublicPriceHistoryResponseDataHistory](docs/PublicPriceHistoryResponseDataHistory.md)
 - [PublicPriceHistoryResponseDataPriceHistory](docs/PublicPriceHistoryResponseDataPriceHistory.md)
 - [PublicSymbolsResponse](docs/PublicSymbolsResponse.md)
 - [PublicSymbolsResponseData](docs/PublicSymbolsResponseData.md)
 - [PublicSymbolsResponseDataSymbols](docs/PublicSymbolsResponseDataSymbols.md)
 - [PublicTrendResponse](docs/PublicTrendResponse.md)
 - [PublicTrendResponseData](docs/PublicTrendResponseData.md)
 - [PublicTrendResponseDataTrend](docs/PublicTrendResponseDataTrend.md)
 - [SymbolsRoute](docs/SymbolsRoute.md)
 - [TrendRoute](docs/TrendRoute.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



