package config

type ServerConfig struct {
	AppVersion    string
	RequestLog    bool
	ListenAddress string
	CorsAllowAll  bool
}
