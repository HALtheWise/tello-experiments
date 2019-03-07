package vins_test

import (
	"fmt"
	"testing"

	"github.com/HALtheWise/tello-experiments/gofly/vins"
)

func TestStuff(t *testing.T) {
	// t.Logf("Got bool: %v", vins.GetBool())
	e := vins.NewEstimator()
	fmt.Printf("%v", e)
}
