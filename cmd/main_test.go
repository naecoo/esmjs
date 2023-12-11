package cmd

import (
	"os"
	"testing"
)

func TestBuild(t *testing.T) {
	output := "/tmp/esmjs_output.js"
	input := "../example/normal/input.js"
	defer os.Remove(output)
	Build(input, output, true, false)
}

func BenchmarkBuild(b *testing.B) {
	output := "/tmp/esmjs_output.js"
	input := "../example/normal/input.js"
	defer os.Remove(output)
	for n := 0; n < b.N; n++ {
		Build(input, output, true, false)
	}
}
