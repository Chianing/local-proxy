package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"local-proxy/constants"
	"local-proxy/model"
	"net/http"
	"os"
	"strings"
)

const configUniqKeySplit = ":"

func InitServer(coreConfig *model.CoreConfig) error {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	engine.Any("/*path", func(ctx *gin.Context) {
		reqPath := ctx.Request.URL.Path

		// reload
		if reqPath == "/-/reload" {
			err := ReloadProxyConfig(coreConfig)
			if err != nil {
				ctx.JSON(int(constants.CodeFailure), model.BuildFailureResponse(fmt.Sprintf("reload proxy config error: %s", err)))
			} else {
				ctx.JSON(int(constants.CodeSuccess), model.BuildSuccessResponse(nil))
			}
			return
		}

		// list
		if reqPath == "/-/list" {
			ctx.JSON(int(constants.CodeSuccess), model.BuildSuccessResponse(*coreConfig.ProxyConfigContentMap))
			return
		}

		// dynamic
		doDynamicProxy(ctx, coreConfig, reqPath)

	})

	return engine.Run(fmt.Sprintf(":%v", *coreConfig.ListenPort))
}

func doDynamicProxy(ctx *gin.Context, coreConfig *model.CoreConfig, reqPath string) {
	// check proxy config exists
	method := ctx.Request.Method
	config, ok := (*coreConfig.ProxyConfigContentMap)[method+configUniqKeySplit+reqPath]
	if !ok {
		// deal restful
		splitReqPath := strings.Split(reqPath, "/")
		restfulPath := strings.Join(splitReqPath[0:len(splitReqPath)-1], "/")

		config, ok = (*coreConfig.ProxyConfigContentMap)[method+configUniqKeySplit+restfulPath]
		if !ok {
			ctx.JSON(int(constants.CodeFailure), model.BuildFailureResponse(fmt.Sprintf("no match mock config for url %s, method: %s", reqPath, method)))
			return
		}
	}

	// return mock result
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(config.MockResult), &res)
	if err != nil {
		ctx.JSON(int(constants.CodeFailure), model.BuildFailureResponse(fmt.Sprintf("unmarshal mock result for url %s error: %s", reqPath, err)))
		return
	}

	ctx.JSON(int(constants.CodeSuccess), res)
}

func ReloadProxyConfig(coreConfig *model.CoreConfig) error {
	fileBytes, err := os.ReadFile(*coreConfig.ProxyConfigPath)
	if err != nil {
		return err
	}

	configTmp := &[]model.ProxyConfig{}
	configMapTmp := make(map[string]model.ProxyConfig)

	err = json.Unmarshal(fileBytes, configTmp)
	if err != nil {
		return err
	}

	for _, config := range *configTmp {
		url := strings.TrimRight(strings.TrimSpace(config.MockUrl), "/")
		method := strings.TrimSpace(config.RequestMethod)
		if method == "" {
			method = http.MethodGet
		}

		configMapTmp[strings.ToUpper(method)+configUniqKeySplit+url] = config
	}

	coreConfig.ProxyConfigContentMap = &configMapTmp

	return nil
}
