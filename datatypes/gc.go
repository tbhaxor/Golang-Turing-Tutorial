package main

import (
	"fmt"
	"runtime"
)

const maxHeapSize = 2 * 1024 * 1024 * 1024 // 1GB

func runGC() {
	fmt.Println("--- Demonstrating GC ---")
	// create 2 gb of junk memory
	var junk = make([]byte, maxHeapSize)
	fmt.Println(junk[:10])

	// run gc manually on the allocated
	fmt.Println("Running GC")
	runtime.GC()
	fmt.Println("\nGC Finished")
}
