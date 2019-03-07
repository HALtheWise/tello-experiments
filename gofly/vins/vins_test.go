package vins_test

import "testing"
import "github.com/HALtheWise/tello-experiments/gofly/vins"

func TestStuff(t *testing.T) {
	t.Logf("Got bool: %v", vins.GetBool())
	var _ vins.Estimator
}
