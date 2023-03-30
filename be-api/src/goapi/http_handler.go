package goapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/btnguyen2k/docms/be-api/src/itineris"
	"github.com/labstack/echo/v4"
)

var (
	// mapping uri -> {http_method -> api_name}
	httpRoutingMap        = map[string]map[string]string{}
	httpHeaderAppId       string
	httpHeaderAccessToken string
)

func registerHttpHandler(uri, httpMethod, apiName string) {
	_, ok := httpRoutingMap[uri]
	if !ok {
		httpRoutingMap[uri] = map[string]string{}
	}
	httpRoutingMap[uri][strings.ToUpper(httpMethod)] = apiName
}

func _parseRequest(apiName string, c echo.Context) (*itineris.ApiContext, *itineris.ApiAuth, *itineris.ApiParams) {
	httpMethod := c.Request().Method
	ctx := itineris.NewApiContext().SetApiName(apiName).SetGateway("HTTP").
		SetContextValue(itineris.CtxHttpMethod, httpMethod).
		SetContextValue(itineris.CtxHttpRequestUrl, c.Request().URL.String()).
		SetContextValue(itineris.CtxClientRealAddr, c.RealIP()).
		SetContextValue(itineris.CtxClientAddr, c.Request().RemoteAddr)
	for k, _ := range c.Request().Header {
		if k != httpHeaderAppId && k != httpHeaderAccessToken {
			k = strings.TrimSpace(strings.ToLower(k))
			if strings.HasPrefix(k, "x-") {
				ctx.SetContextValue(k[2:], c.Request().Header.Get(k))
			}
		}
	}

	auth := itineris.NewApiAuth(c.Request().Header.Get(httpHeaderAppId), c.Request().Header.Get(httpHeaderAccessToken))

	params := itineris.NewApiParams()
	// first, populate params passed via request body
	if !strings.EqualFold("GET", httpMethod) && !strings.EqualFold("HEAD", httpMethod) {
		requestBodyData := map[string]interface{}{}
		if err := c.Bind(&requestBodyData); err != nil {
			log.Printf("Error while parsing request body as Json: " + err.Error())
			log.Printf("Request: " + ctx.ToJsonString())
		} else {
			for k, v := range requestBodyData {
				params.SetParam(k, v)
			}
		}
	}
	// second, populate params on URI path
	for _, p := range c.ParamNames() {
		params.SetParam(p, c.Param(p))
	}

	// finally, populate params on query string
	for k, v := range c.QueryParams() {
		if v != nil && len(v) > 0 {
			params.SetParam(k, v[0])
		}
	}

	return ctx, auth, params
}

// Handle API request via HTTP.
func apiHttpHandler(c echo.Context) error {
	uriPattern := c.Path()
	if _, ok := httpRoutingMap[uriPattern]; !ok {
		return c.JSON(http.StatusOK, itineris.ResultNotImplemented.ToMap())
	}
	httpMethod := strings.ToUpper(c.Request().Method)
	if _, ok := httpRoutingMap[uriPattern][httpMethod]; !ok {
		return c.JSON(http.StatusOK, itineris.ResultNotImplemented.ToMap())
	}

	apiName := httpRoutingMap[uriPattern][httpMethod]
	ctx, auth, params := _parseRequest(apiName, c)

	apiResult := ApiRouter.CallApi(ctx, auth, params)
	return c.JSON(http.StatusOK, apiResult.ToMap())
}
