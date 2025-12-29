package main


type Config struct {
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	Port        int    `json:"port"        yaml:"port"`
	Replicas    int    `json:"replicas"    yaml:"replicas"`
}


type ValidationResult struct {
	Success bool
	Message string
}

type Rule interface {
	Validate(c Config) (bool , string)
}