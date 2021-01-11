package gotools

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGetSeedForRandomCreation(t *testing.T) {

	seed := GetSeedForRandomCreation()
	t.Log(seed)
	if seed == 0 {
		t.Error("failed to create seed by host address.")
	}

}

func TestExampleForRandomDelaying(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(GetSeedForRandomCreation())
	for i := 0; i < 10; i++ {
		randomDelay := rand.Intn(100)
		fmt.Println(randomDelay)
	}
}
