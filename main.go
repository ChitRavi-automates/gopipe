package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// A Step is one unit of work in the pipeline.
type Step struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

func main() {
	data, err := os.ReadFile("pipeline.yaml")
	if err != nil {
		fmt.Println("could not read pipeline.yaml:", err)
		os.Exit(1)
	}

	var pipeline Pipeline
	err = yaml.Unmarshal(data, &pipeline)
	if err != nil {
		fmt.Println("could not parse pipeline.yaml:", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	failed := false

	for _, step := range pipeline.Steps {
		wg.Add(1)

		go func(step Step) {
			defer wg.Done()

			start := time.Now()
			cmd := exec.Command("bash", "-c", step.Command)
			output, err := cmd.CombinedOutput()
			duration := time.Since(start)

			result := fmt.Sprintf("▶ %s\n%s", step.Name, string(output))
			if err != nil {
				result += fmt.Sprintf("✗ %s (%s)\n", step.Name, duration)
				mu.Lock()
				failed = true
				mu.Unlock()
			} else {
				result += fmt.Sprintf("✓ %s (%s)\n", step.Name, duration)
			}
			fmt.Print(result)
		}(step)
	}

	wg.Wait()

	if failed {
		os.Exit(1)
	}
}
