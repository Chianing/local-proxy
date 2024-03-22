package main

import (
	"flag"
	"fmt"
	"local-proxy/model"
	"local-proxy/server"
)

var coreConfig *model.CoreConfig

func initConfig() error {
	coreConfig = &model.CoreConfig{
		ProxyConfigPath: flag.String("config.path", "./proxy-config.json", "proxy config path, should be json file"),
		ListenPort:      flag.Int("listen.port", 8080, "server listen port"),
	}

	flag.Parse()

	err := server.ReloadProxyConfig(coreConfig)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initConfig()
	if err != nil {
		panic(fmt.Sprintf("init config error: %s", err))
	}

	err = server.InitServer(coreConfig)
	if err != nil {
		panic(fmt.Sprintf("init server error: %s", err))
	}
}
