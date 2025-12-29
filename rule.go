package main

import "fmt"


type PortRule struct {}
func (p PortRule) Validate(c Config) (bool , string) {
	if c.Port <= 0 {
		return false , "The Port Number should be greater than 0"
	}
	return true ,"Port check passed"
}

type ReplicaRule struct{
	MinReplicas int
}

func (r ReplicaRule) Validate(c Config) (bool, string) {
	if c.Replicas < r.MinReplicas {
		return false, fmt.Sprintf("Replicas %d is below minimum %d", c.Replicas, r.MinReplicas)
	}
	return true, "Replica Check Passed"
}
