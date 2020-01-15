/*
 * CryptoWeather
 *
 * The CryptoWeather API allows for access to all of the cryptocurrency data and market forecast services provided. There are two primary categories of routes, `public` and `private`, where `public` routes are accessible to the general public and do not require API authentication, and `private` routes, which require API authentication.  ## General Overview  1. All API methods adhere to RESTful best practices as closely as possible. As such, all API calls will be made via the standard HTTP protocol using the GET/POST/PUT/DELETE request types.  2. Every request returns the status as a JSON response with the following:     - success, true if it was successful     - code, the http status code (also in the response header)         - 200 if response is successful         - 400 if bad request         - 401 if authorization JWT is wrong or limit exceeded         - 404 wrong route         - 500 for any internal errors     - status, the status of the request, default **success**     - errors, an array of any relevant error details  3. For any requests that are not successful an error message is specified and returned as an array for the **errors** key in the JSON response.  4. All authentication uses JSON Web Tokens (JWT), which is set as the **Authorization** entry in the header, see the following for more details.     - http://jwt.io     - https://scotch.io/tutorials/the-anatomy-of-a-json-web-token  ## Code Example  The following is a code example in Python, which demonstrates using the [Python Requests library](https://requests.readthedocs.io/en/master/) for both the `public` and `private` API routes.  ``` import requests  HOST = \"https://api.cryptoweather.ai/v1\"  # Your API key (JWT) API_KEY = \"<YOUR API KEY>\"  # Example public request, no API key required. requests.get(\"{}/public/symbols\".format(HOST)).json()  # Get the current btc price using the public route requests.get(\"{}/public/price-current/{}\".format(HOST, \"btc\")).json()   # Example private request, API key required. Get the btc hourly forecasts headers = {\"Authorization\": \"Bearer {}\".format(API_KEY)} requests.get(\"{}/private/forecast/{}/{}\".format(HOST, \"btc\", \"1h\"),              headers=headers).json() ```
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type PrivateApiService service

/* 
PrivateApiService Accuracy
The accuracy response contains the following attributes.  + &#x60;rmse&#x60; Root Mean Square Error  + &#x60;mae&#x60; Mean Absolute error  + &#x60;r2&#x60; R Squared  + &#x60;ci&#x60; Confidence Interval lower and upper error bounds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol The cryptocurrency symbol.
 * @param interval The forecast interval, 1h or 1d.
 * @param period The period for computing the accuracy, such as the past 7 days.
 * @param optional nil or *V1PrivateAccuracySymbolIntervalPeriodGetOpts - Optional Parameters:
     * @param "Cookie" (optional.String) -  e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89
     * @param "XCsrf" (optional.String) -  e.g. b1820141-1bad-4a9c-93c0-52022817ce89

@return PrivateAccuracyResponse
*/

type V1PrivateAccuracySymbolIntervalPeriodGetOpts struct { 
	Cookie optional.String
	XCsrf optional.String
}

func (a *PrivateApiService) V1PrivateAccuracySymbolIntervalPeriodGet(ctx context.Context, symbol string, interval string, period string, localVarOptionals *V1PrivateAccuracySymbolIntervalPeriodGetOpts) (PrivateAccuracyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PrivateAccuracyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/private/accuracy/{symbol}/{interval}/{period}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", fmt.Sprintf("%v", symbol), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interval"+"}", fmt.Sprintf("%v", interval), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"period"+"}", fmt.Sprintf("%v", period), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Cookie.IsSet() {
		localVarHeaderParams["Cookie"] = parameterToString(localVarOptionals.Cookie.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XCsrf.IsSet() {
		localVarHeaderParams["X-csrf"] = parameterToString(localVarOptionals.XCsrf.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v PrivateAccuracyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 400 {
			var v DefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
PrivateApiService Forecast
The forecast response contains a sequence of forecasts at the specified intervals, with the following attributes.  + &#x60;time_start&#x60; start time of the period the forecast is applicable for  + &#x60;time_end&#x60; end time of the period the forecast is applicable for  + &#x60;low&#x60; forecasted high during the period  + &#x60;high&#x60; forecasted low during the period  + &#x60;weighted_price&#x60; forecasted weighted price during the period  + &#x60;change_pct&#x60; percent change in price for forecasted weighted_price relative to current price  + &#x60;change_usd&#x60; dollar change in price for forecasted weighted_price relative to current price
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol The cryptocurrency symbol.
 * @param interval The forecast interval, 1h or 1d.
 * @param optional nil or *V1PrivateForecastSymbolIntervalGetOpts - Optional Parameters:
     * @param "Cookie" (optional.String) -  e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89
     * @param "XCsrf" (optional.String) -  e.g. b1820141-1bad-4a9c-93c0-52022817ce89

@return PrivateForecastResponse
*/

type V1PrivateForecastSymbolIntervalGetOpts struct { 
	Cookie optional.String
	XCsrf optional.String
}

func (a *PrivateApiService) V1PrivateForecastSymbolIntervalGet(ctx context.Context, symbol string, interval string, localVarOptionals *V1PrivateForecastSymbolIntervalGetOpts) (PrivateForecastResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PrivateForecastResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/private/forecast/{symbol}/{interval}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", fmt.Sprintf("%v", symbol), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interval"+"}", fmt.Sprintf("%v", interval), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Cookie.IsSet() {
		localVarHeaderParams["Cookie"] = parameterToString(localVarOptionals.Cookie.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XCsrf.IsSet() {
		localVarHeaderParams["X-csrf"] = parameterToString(localVarOptionals.XCsrf.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v PrivateForecastResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 400 {
			var v DefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
PrivateApiService Trend
The trend response contains a collection of forecasts for different intervals with the following attributes.  + &#x60;time_start&#x60; start time of the period the forecast is applicable for  + &#x60;time_end&#x60; end time of the period the forecast is applicable for  + &#x60;interval&#x60; interval in hours that the forecast is applicable for  + &#x60;weighted_price&#x60; forecasted weighted price during the period  + &#x60;change_pct&#x60; percent change in price for forecasted weighted_price relative to current price  + &#x60;change_usd&#x60; dollar change in price for forecasted weighted_price relative to current price
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param symbol The cryptocurrency symbol.
 * @param optional nil or *V1PrivateTrendSymbolGetOpts - Optional Parameters:
     * @param "Cookie" (optional.String) -  e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89
     * @param "XCsrf" (optional.String) -  e.g. b1820141-1bad-4a9c-93c0-52022817ce89

@return PublicTrendResponse
*/

type V1PrivateTrendSymbolGetOpts struct { 
	Cookie optional.String
	XCsrf optional.String
}

func (a *PrivateApiService) V1PrivateTrendSymbolGet(ctx context.Context, symbol string, localVarOptionals *V1PrivateTrendSymbolGetOpts) (PublicTrendResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PublicTrendResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/private/trend/{symbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"symbol"+"}", fmt.Sprintf("%v", symbol), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Cookie.IsSet() {
		localVarHeaderParams["Cookie"] = parameterToString(localVarOptionals.Cookie.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XCsrf.IsSet() {
		localVarHeaderParams["X-csrf"] = parameterToString(localVarOptionals.XCsrf.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v PublicTrendResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 400 {
			var v DefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
PrivateApiService Trend Tabular
The trend tabular response contains a collection of forecasts for all supported cryptocurrencies at different intervals with the following attributes.  + &#x60;time_start&#x60; start time of the period the forecast is applicable for  + &#x60;time_end&#x60; end time of the period the forecast is applicable for  + &#x60;interval&#x60; interval in hours that the forecast is applicable for  + &#x60;weighted_price&#x60; forecasted weighted price during the period  + &#x60;change_pct&#x60; percent change in price for forecasted weighted_price relative to current price  + &#x60;change_usd&#x60; dollar change in price for forecasted weighted_price relative to current price
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *V1PrivateTrendTabularGetOpts - Optional Parameters:
     * @param "Cookie" (optional.String) -  e.g. csrf&#x3D;b1820141-1bad-4a9c-93c0-52022817ce89
     * @param "XCsrf" (optional.String) -  e.g. b1820141-1bad-4a9c-93c0-52022817ce89

@return PrivateTrendTabularResponse
*/

type V1PrivateTrendTabularGetOpts struct { 
	Cookie optional.String
	XCsrf optional.String
}

func (a *PrivateApiService) V1PrivateTrendTabularGet(ctx context.Context, localVarOptionals *V1PrivateTrendTabularGetOpts) (PrivateTrendTabularResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PrivateTrendTabularResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/private/trend-tabular"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Cookie.IsSet() {
		localVarHeaderParams["Cookie"] = parameterToString(localVarOptionals.Cookie.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XCsrf.IsSet() {
		localVarHeaderParams["X-csrf"] = parameterToString(localVarOptionals.XCsrf.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v PrivateTrendTabularResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 400 {
			var v DefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}