package warp

import "testing"

func TestWarp(t *testing.T) {
	src := []float64{1.3, 1.2, 1.5, 1.4, 1.4, 1.1, 1.4}
	tgt := []float64{2.2, 2.3, 2.6, 2.4, 2.3, 2.0, 2.3}

	var warp1, warp2 float64
	{
		warp1, _ = DTW(src, tgt, L1)
		warp2, _ = WindowDTW(src, tgt, L2, 2.0)
		t.Logf("\nL1: %v \nL2: %v", warp1, warp2)
	}
	{
		warp1, _ = DTW(src, tgt, L1)
		warp2, _ = WindowDTW(src, tgt, L2, 2.0)
		t.Logf("\nL1: %v \nL2: %v", warp1, warp2)
	}
}
