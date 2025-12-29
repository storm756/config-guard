# ConfigGuard 🛡️

A high-performance CLI tool for validating infrastructure configuration files, built with **Golang**.

## 🚀 Why this project?
I built this to understand the underlying architecture of policy engines like **Kyverno**. It mimics the behavior of an admission controller by scanning configuration files (**JSON & YAML**) and validating them against a set of extensible rules.

## 🛠️ Tech Stack & Key Concepts
*   **Golang**: Core logic and type safety.
*   **Concurrency**: Uses `Goroutines` and `Channels` to process multiple files in parallel (high throughput).
*   **Interfaces**: Implements the `Rule` pattern (Strategy Design Pattern) to make the validator easily extensible.
*   **Sync**: Uses `sync.WaitGroup` to manage concurrent execution flow.
*   **Parsing**: Supports both standard JSON (`encoding/json`) and YAML (`gopkg.in/yaml.v3`) to match Kubernetes standards.

## ⚡ How it works
The application reads a directory of configuration files and spins up a lightweight thread (Goroutine) for each file. Results are aggregated via a buffered Channel to ensure thread-safe reporting.

```bash
# Run the validator
go run . -dir=./test_configs