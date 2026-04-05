package test

import (
	"cli-text-compare/pkg"
	"testing"
	"time"
)

func TestPerformanceOfDiff(t *testing.T) {

	textOne := readFile(t, "./perf_one.json")
	textTwo := readFile(t, "./perf_two.json")

	start := time.Now()
	actions := pkg.Diff(textOne, textTwo, false)
	elapsed := time.Since(start)
	t.Logf("Diff took %v ms", elapsed.Milliseconds())

	if len(actions) == 0 {
		t.Log("somehow the actions are incorrect")
		t.FailNow()
	}

	// I feel like anything bigger than a second and I'd go, fuuuucking hell and delete the app
	// so, that's kind of a cap
	// and for ~2.5k lines of json vs 3.5k lines of json, 1 second is perfectly fine I think
	if elapsed > 1*time.Second {
		t.Log("if the accuracy is okay, go for it but ya' fucked the performance of the thing")
		t.FailNow()
	}
}
