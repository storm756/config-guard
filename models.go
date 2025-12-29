package main


type Config struct {
	ServiceName string `json:"serviceName"`
	Port int `json:"port"`
	Replicas int `json:"replicas"`
	Enabled bool `json:"enabled"`
}


type ValidationResult struct {
	Success bool
	Message string
}

type Rule interface {
	Validate(c Config) (bool , string)
}