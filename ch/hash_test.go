package ch

import (
	"time"
	"testing"
	"math/rand"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSomeHashValues(t *testing.T) {

	gorizon := Murmur3([]byte("gorizon"))
	ugol := Murmur3([]byte("ugol"))

	gorizon_expected := uint32(1974954250)
	ugol_expected := uint32(112532158)

	if gorizon != gorizon_expected {
		t.Errorf("For 'gorizon' expecting '%d', but got '%d'", gorizon_expected, gorizon)
	}

	if ugol != ugol_expected {
		t.Errorf("For 'ugol' expecting  '%d', but got '%d'", ugol_expected, ugol)
	}
}
