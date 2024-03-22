package model

type CoreConfig struct {
	ListenPort            *int
	ProxyConfigPath       *string
	ProxyConfigContentMap *map[string]ProxyConfig
}

type ProxyConfig struct {
	MockUrl       string `json:"mockUrl"`
	RequestMethod string `json:"requestMethod"`
	MockResult    string `json:"mockResult"`
}
