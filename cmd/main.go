package main

import (
	"sync"

	alg "github.com/adi221/playground/advanced-concurrency-patterns/pkg/algorithms"
	con "github.com/adi221/playground/advanced-concurrency-patterns/pkg/concepts"
	gen "github.com/adi221/playground/advanced-concurrency-patterns/pkg/generators"
	pri "github.com/adi221/playground/advanced-concurrency-patterns/pkg/primitives"
)

type runner []func()

type Runner struct {
	runner
}

// This function executes all Runner's operations concurrently.
func (r *Runner) Run() {

	var wg sync.WaitGroup
	wg.Add(len((*r).runner))

	for _, fn := range (*r).runner {
		go func(wg *sync.WaitGroup, fn func()) {
			defer (*wg).Done()
			fn()
		}(&wg, fn)
	}

	wg.Wait()
}

var AlgorithmsRunner Runner = Runner{[]func(){alg.BinarySearch}}
var ConceptsRunner Runner = Runner{
	[]func(){
		con.OrChannel,
		con.LexicalConfinment,
		con.NoErrorHandling,
		con.ErrorHandling,
	},
}
var PrimitivesRunner Runner = Runner{
	[]func(){
		pri.Cond,
		pri.RWMutex,
		pri.Broadcast,
		pri.KnightsTour,
	},
}

func RunAll() {
	ConceptsRunner.Run()
	PrimitivesRunner.Run()
	AlgorithmsRunner.Run()
}

func main() {
	gen.Generators()
	RunAll()
}
