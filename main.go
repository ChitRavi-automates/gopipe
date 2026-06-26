package main

import (
	"fmt"
	"os"
	"os/exec"
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

	failed := false

	for _, step := range pipeline.Steps {
		fmt.Printf("▶ %s\n", step.Name) 

		start := time.Now()

		cmd := exec.Command("bash", "-c", step.Command)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))

		duration := time.Since(start)

		if err != nil {
			fmt.Printf("✗ %s (%s)\n", step.Name, duration)
			failed = true
		} else {
			fmt.Printf("✓ %s (%s)\n", step.Name, duration)
		}
	}

	if failed {
		os.Exit(1)
	}
}
