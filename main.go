package main

import (
	"fmt"
	"os"
	"os/exec"
)

// A Step is one unit of work in the pipeline.
type Step struct {
	Name    string
	Command string // the shell command to run, e.g. "go version"
}

func main() {
	// 1. Define a few steps to run.
	steps := []Step{
		{Name: "Go version", Command: "go version"},
		{Name: "List files", Command: "ls -la"},
		{Name: "Say hello", Command: "echo hello from gopipe"},
                {Name: "Broken step", Command: "exit 1"},
	}


	failed := false

	// 2. Loop over each step and run it.
	for _, step := range steps {
		fmt.Printf("▶ %s\n", step.Name)

		cmd := exec.Command("bash", "-c", step.Command)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))

		if err != nil {
			fmt.Printf("✗ %s\n", step.Name)
			failed = true
		} else {
			fmt.Printf("✓ %s\n", step.Name)
		}
	}

	// 3. Exit non-zero if anything failed, so CI can detect failure.
	if failed {
		os.Exit(1)
	}
}
