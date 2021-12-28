package mascot_test

import (
	"testing"

	"github.com/nomad31415/goQuantum/mascot"
)

func TestMascot(t *testing.T) {

	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Bad string")
	}
}
