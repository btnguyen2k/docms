package goapi

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/btnguyen2k/docms/be-api/src/itineris"
	"github.com/btnguyen2k/docms/be-api/src/utils"
	hocon "github.com/go-akka/configuration"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultConfigFile = "./config/application.conf"
	logLevelInfo      = "INFO"
	logLevelWarning   = "WARN"
	logLevelError     = "ERROR"
)

var (
	AppVersion        string
	AppVersionNumber  uint64
	AppConfig         *hocon.Config
	ApiRouter         *itineris.ApiRouter
	PostInitEchoSetup = make([]func(e *echo.Echo) error, 0)
)

// Start bootstraps the application.
func Start(bootstrappers ...IBootstrapper) {
	var err error

	// load application configurations
	AppConfig = initAppConfig()
	httpHeaderAppId = AppConfig.GetString("api.http.header_app_id")
	httpHeaderAccessToken = AppConfig.GetString("api.http.header_access_token")
	AppVersion = AppConfig.GetString("app.version")
	AppVersionNumber = utils.VersionToNumber(AppVersion)

	// setup api-router
	ApiRouter = itineris.NewApiRouter()

	// initialize "Location"
	utils.Location, err = time.LoadLocation(AppConfig.GetString("timezone"))
	if err != nil {
		panic(err)
	}

	// bootstrapping
	if bootstrappers != nil {
		for _, b := range bootstrappers {
			log.Printf("[%s] Bootstrapping %#v...", logLevelInfo, b)
			err := b.Bootstrap()
			if err != nil {
				log.Println(err)
			}
		}
	}

	// initialize and start echo server
	initEchoServer()
}

func initAppConfig() *hocon.Config {
	const envKey = "APP_CONFIG"
	configFile := os.Getenv(envKey)
	if configFile == "" {
		log.Printf("[%s] No environment %s found, fallback to [%s]", logLevelWarning, envKey, defaultConfigFile)
		configFile = defaultConfigFile
	}
	return loadAppConfig(configFile)
}

func initEchoServer() {
	const confKeyHttpPort = "api.http.listen_port"
	listenPort := AppConfig.GetInt32(confKeyHttpPort, 0)
	if listenPort <= 0 {
		log.Printf("[%s] No valid [%s] configured, REST API gateway is disabled", logLevelWarning, confKeyHttpPort)
		return
	}
	listenAddr := AppConfig.GetString("api.http.listen_addr", "127.0.0.1")
	e := echo.New()
	requestTimeout := AppConfig.GetTimeDuration("api.request_timeout", time.Duration(0))
	if requestTimeout > 0 {
		e.Server.ReadTimeout = requestTimeout
	}
	bodyLimit := AppConfig.GetByteSize("api.max_request_size")
	if bodyLimit != nil && bodyLimit.Int64() > 0 {
		e.Use(middleware.BodyLimit(bodyLimit.String()))
	}
	allowOgirinsStr := AppConfig.GetString("api.http.allow_origins", "*")
	if allowOgirins := regexp.MustCompile("[,;\\s]+").Split(allowOgirinsStr, -1); len(allowOgirins) > 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowOgirins,
		}))
	}

	// register API http endpoints
	const confKeyHttpEndpoints = "api.http.endpoints"
	hasEndpoints := false
	confV := AppConfig.GetValue(confKeyHttpEndpoints)
	if confV != nil && confV.IsObject() {
		for uri, uriO := range confV.GetObject().Items() {
			if uriO.IsObject() && !uriO.IsEmpty() {
				hasEndpoints = true
				e.Any(uri, apiHttpHandler)
				for httpMethod, apiName := range uriO.GetObject().Items() {
					registerHttpHandler(uri, httpMethod, apiName.GetString())
				}
			}
		}
	}
	js, _ := json.Marshal(httpRoutingMap)
	log.Printf("[%s] API http endpoints: %s", logLevelInfo, js)
	if !hasEndpoints {
		log.Printf("[%s] No valid HTTP endpoints defined at key [%s]", logLevelWarning, confKeyHttpEndpoints)
	}

	for _, postInitEchoSetup := range PostInitEchoSetup {
		if err := postInitEchoSetup(e); err != nil {
			log.Printf("[%s] Error executing post-init Echo setup:  %s", logLevelError, err)
		}
	}

	log.Printf("[%s] Starting [%s] RESTful server on [%s:%d]...",
		logLevelInfo,
		AppConfig.GetString("app.name")+" v"+AppConfig.GetString("app.version"),
		listenAddr, listenPort)
	go e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", listenAddr, listenPort)))
}
