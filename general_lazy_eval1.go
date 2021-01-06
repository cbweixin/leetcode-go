package main

import "fmt"

type Any interface{}

type EvalFunc func(Any) (Any, Any)

func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}

	evenFunc2 := func(state Any) (Any, Any) {
		os := state.(string)
		ns := os + "a"
		return os, ns
	}

	s := BuildLazyIntEvaluator2(evenFunc2, "b")
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, s())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func BuildLazyIntEvaluator2(evalFunc EvalFunc, initState Any) func() string {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() string {
		return ef().(string)
	}
}
