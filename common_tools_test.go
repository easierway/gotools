package gotools

import "testing"

func TestGetSeedForRandomCreation(t *testing.T) {

	seed := GetSeedForRandomCreation()
	t.Log(seed)
	if seed == 0 {
		t.Error("failed to create seed by host address.")
	}
}
