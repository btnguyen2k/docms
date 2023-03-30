package gvabe

import (
	"crypto/x509"
	"encoding/pem"
	"runtime"
	"strconv"

	"github.com/btnguyen2k/docms/be-api/src/goapi"
	"github.com/btnguyen2k/docms/be-api/src/itineris"
)

// Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)
//   - api-handler function must have the following signature:
//     func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
func initApiHandlers(router *itineris.ApiRouter) {
	router.SetHandler("health", apiHealth)
	router.SetHandler("info", apiInfo)
}

/*------------------------------ shared variables and functions ------------------------------*/

var (
	// those APIs will not need authentication.
	// "false" means client, however, needs to send app-id along with the API call
	// "true" means the API is free for public call
	publicApis = map[string]bool{
		"info":   true,
		"health": true,
	}
)

/*------------------------------ APIs ------------------------------*/

var _apiResultHealth = itineris.NewApiResult(itineris.StatusOk).SetData("ok")

// API handler "health"
func apiHealth(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	return _apiResultHealth
}

// API handler "info"
func apiInfo(_ *itineris.ApiContext, _ *itineris.ApiAuth, _ *itineris.ApiParams) *itineris.ApiResult {
	var publicPEM []byte
	if pubDER, err := x509.MarshalPKIXPublicKey(rsaPubKey); err == nil {
		pubBlock := pem.Block{
			Type:    "PUBLIC KEY",
			Headers: nil,
			Bytes:   pubDER,
		}
		publicPEM = pem.EncodeToMemory(&pubBlock)
	} else {
		publicPEM = []byte(err.Error())
	}

	result := map[string]interface{}{
		"app": map[string]interface{}{
			"name":        goapi.AppConfig.GetString("app.name"),
			"shortname":   goapi.AppConfig.GetString("app.shortname"),
			"version":     goapi.AppConfig.GetString("app.version"),
			"description": goapi.AppConfig.GetString("app.desc"),
		},
		"exter": map[string]interface{}{
			"app_id":   exterAppId,
			"base_url": exterBaseUrl,
		},
		"rsa_public_key": string(publicPEM),
		"debug_mode":     DEBUG_MODE,
		"demo_mode":      DEMO_MODE,
	}
	if DEMO_MODE {
		result["demo"] = map[string]interface{}{
			"user_id":   goapi.AppConfig.GetString("gvabe.init.admin_user_id"),
			"user_pwd":  goapi.AppConfig.GetString("gvabe.init.admin_user_pwd"),
			"user_name": goapi.AppConfig.GetString("gvabe.init.admin_user_name"),
		}
	}
	if DEBUG_MODE || DEMO_MODE {
		/*!!! demo purpose only! exposing memory usage is generally not a good idea !!!*/
		var m runtime.MemStats
		result["memory"] = map[string]interface{}{
			"alloc":     m.Alloc,
			"alloc_str": strconv.FormatFloat(float64(m.Alloc)/1024.0/1024.0, 'f', 1, 64) + " MiB",
			"sys":       m.Sys,
			"sys_str":   strconv.FormatFloat(float64(m.Sys)/1024.0/1024.0, 'f', 1, 64) + " MiB",
			"gc":        m.NumGC,
		}
	}
	return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}
