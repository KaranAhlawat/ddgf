package server

type APIConfig struct {
	port string
}

func NewAPIConfig() *APIConfig {
	return &APIConfig{port: ":3000"}
}
