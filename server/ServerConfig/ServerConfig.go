package ServerConfig

type ServerConfig struct {
	Port int
	Host string
}

func New(port int, host string) *ServerConfig {
	return &ServerConfig{port, host}
}
