package main

import (
	"github.com/vikash/learning-golang/003-concurrency/functions"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	years := []functions.Year{2012, 2013, 2014, 2015}
	for _, y := range years {
		wg.Add(1)

		go func(y functions.Year) {
			functions.WomenRelatedCrimeStats(y)
			wg.Done()
		}(y)
	}

	wg.Wait()

}
