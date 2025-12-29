package main 

import "testing"

func TestPortRule(t *testing.T) {
	rule := PortRule{}


validConfig := Config{Port :8080}
valid,_ := rule.Validate(validConfig)
if !valid {
	t.Error("Expected port 8080 to be valid, but got invalid")
}
invalidConfig := Config{Port :-5}
valid,_ = rule.Validate(invalidConfig)


if valid {
	t.Error("Expected port -5 tp be invalid, but got valid")
}
}

func TestReplicaRule(t *testing.T) {
	// We set minimum replicas to 2
	rule := ReplicaRule{MinReplicas: 2}

	validConfig := Config{Replicas: 3}
	valid, _ := rule.Validate(validConfig)
	if !valid {
		t.Errorf("Expected 3 replicas to pass, but it failed")
	}

	invalidConfig := Config{Replicas: 1}
	valid, _ = rule.Validate(invalidConfig)
	if valid {
		t.Errorf("Expected 1 replica to fail, but it passed")
	}
}