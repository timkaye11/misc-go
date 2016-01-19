package rao

import "testing"

func init() {
	SetModulus(24.0)
}

func TestRao(t *testing.T) {
	obs := []float64{1.0, 5.0, 9.0, 14.0, 20.0}
	alpha := 0.05
	s := Statistic(obs, alpha)
	t.Log(s)
}
