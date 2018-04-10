package goosli

import (
	"runtime"
	"sync"
	"math"
	"strconv"
)

func DoInParallelAndWait(work func(wi, wn int)) {
	wn := runtime.NumCPU()
	var wg sync.WaitGroup
	for wi := 0; wi < wn; wi++ {
		wg.Add(1)
		go func(wi, wn int) {
			work(wi, wn)
			wg.Done()
		}(wi, wn)
	}
	wg.Wait()
}

func DoInParallel(work func(wi, wn int)) {
	wn := runtime.NumCPU()
	for wi := 0; wi < wn; wi++ {
		go func(wi, wn int) {
			work(wi, wn)
		}(wi, wn)
	}
}

func RoundPlaces(a float64, places int) float64 {
	shift := powersOfTen[places]
	return float64(Round(a*shift)) / shift
}

func Round(a float64) int {
	if a < 0 {
		return int(math.Ceil(a - 0.5))
	} else {
		return int(math.Floor(a + 0.5))
	}
}

var powersOfTen = []float64{1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16}

func StrF(f float64) string {
	return strconv.FormatFloat(f, 'f', 3, 64)
}
