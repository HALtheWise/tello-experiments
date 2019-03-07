package vins_test

import (
	"fmt"
	"testing"

	"github.com/HALtheWise/tello-experiments/gofly/vins"
)

func TestStuff(t *testing.T) {
	e := vins.NewEstimator()
	fmt.Printf("Estimator: %v\n", e)
	// fmt.Printf("Initial structure: %v\n", e.InitialStructure())
}
