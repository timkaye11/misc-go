package warp

import "math"

// Define our distance function
type Measure func(float64, float64) float64

var (
	// Median
	L1 Measure = func(x, y float64) float64 {
		return math.Abs(x - y)
	}

	// Mean
	L2 Measure = func(x, y float64) float64 {
		return math.Pow(x-y, 2.0)
	}
)

func WindowDTW(s, t []float64, d Measure, window float64) (float64, [][]float64) {
	return dtw(s, t, d, window)
}

func DTW(s, t []float64, d Measure) (float64, [][]float64) {
	return dtw(s, t, d, 0.0)
}

func dtw(s, t []float64, d Measure, window float64) (float64, [][]float64) {
	dtw := make([][]float64, len(s)+1)

	var w float64
	var usewindow bool
	if window != 0.0 {
		w = math.Max(window, math.Abs(float64(len(s)-len(t))))
		usewindow = true
	}

	for i := 0; i <= len(s); i++ {
		dtw[i] = make([]float64, len(t)+1)
		dtw[i][0] = math.Inf(1)
	}

	for i := 0; i <= len(t); i++ {
		dtw[0][i] = math.Inf(1)
	}

	if usewindow {
		for i := 0; i <= len(s); i++ {
			for j := 0; j <= len(t); j++ {
				dtw[i][j] = math.Inf(1)
			}
		}
	}
	dtw[0][0] = 0.0

	for i := 1; i <= len(s); i++ {
		lo := 1
		hi := len(t)
		if usewindow {
			lo = int(math.Max(1.0, float64(i)-w))
			hi = int(math.Min(float64(len(t)), float64(i)+w))
		}

		for j := lo; j <= hi; j++ {
			cost := d(s[i-1], t[j-1])
			dtw[i][j] = cost + min(
				dtw[i-1][j],
				dtw[i][j-1],
				dtw[i-1][j-1],
			)
		}
	}

	return dtw[len(s)][len(t)], dtw
}

func min(vals ...float64) float64 {
	if len(vals) == 0 {
		return 0.0
	}
	m := vals[0]
	for _, val := range vals {
		if m > val {
			m = val
		}
	}
	return m
}
