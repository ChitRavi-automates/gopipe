# gopipe

A minimal CI pipeline runner written in Go. gopipe reads a list of steps from a
YAML file, runs them as shell commands, and reports whether each step passed or
failed — exiting non-zero if any step fails, so it can run inside CI.

## Why I built it

I'm learning Go, and I wanted a project that mirrors the kind of work a CI/CD
platform team does. gopipe is a small version of what tools like Jenkins and
GitHub Actions do under the hood: take a pipeline definition and execute it.

## How it works

- Steps are defined in `pipeline.yaml`
- gopipe reads and parses the file, then runs each step's command via `bash`
- Steps run concurrently as goroutines, synchronized with a `sync.WaitGroup`
- Each step is timed and reports `✓` on success or `✗` on failure
- If any step fails, gopipe exits with code 1 (a failed build)

## Usage
gorun .
## Example pipeline

```yaml
steps:
  - name: Go version
    command: go version
  - name: List files
    command: ls -la
  - name: Say hello
    command: echo hello from gopipe
```

gopipe runs each step in the file and reports the result, with timing.

## Project status

Work in progress — built step by step as a Go learning project.

## Roadmap

- [x] Run steps from a YAML pipeline file
- [x] Report pass/fail and exit non-zero on failure
- [x] Time each step
- [x] Run steps concurrently with goroutines
- [ ] GitHub Actions CI workflow
- [ ] Tagged release

## Tech

Go · YAML (`gopkg.in/yaml.v3`) · `os/exec` · `sync`
