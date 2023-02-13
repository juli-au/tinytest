package main

import (
	"fmt"
	"time"

	"foxygo.at/evy/pkg/evaluator"
)

var evySource = `
func fn
  x := 5
  for n := range [1]
    print "n:" n "x:" x
    x = x - 1
  end
end

fn
`

func main() {
	rt := evaluator.Runtime{
		Print: func(s string) { fmt.Print(s) },
	}
	builtins := evaluator.DefaultBuiltins(rt)
	eval := evaluator.NewEvaluator(builtins)
	eval.Yield = newSleepingYielder()
	eval.Run(evySource)
}

// newSleepingYielder yields the CPU so that JavaScript/browser events
// get a chance to be processed. Currently(Feb 2023) it seems that you
// can only yield to JS by sleeping for at least 1ms but having that
// delay is not ideal. Other methods of yielding can be explored by
// implementing a different yield function.
func newSleepingYielder() func() {
	count := 0
	start := time.Now()
	return func() {
		count++
		if count > 1000 && time.Since(start) > 100*time.Millisecond {
			time.Sleep(time.Millisecond)
			start = time.Now()
			count = 0
		}
	}
}
