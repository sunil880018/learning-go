# Go Testing Guide

This document provides a concise overview of the most commonly used testing types in Go, the tools/packages used, and their real-world applications.

| Type                    | Tool / Package           | Real Use Case               |
| ----------------------- | ------------------------ | --------------------------- |
| Unit Test               | `testing`                | Function logic              |
| Mocking                 | `testify/mock`, `gomock` | DB/API dependencies         |
| Integration Test        | `dockertest`, real DB    | Multi-component tests       |
| HTTP / API Test         | `httptest`               | REST / HTTP handlers        |
| Benchmark / Performance | `testing.B`              | Measure speed, optimization |

⚡ **Pro tips:**

- **Unit tests + mocking** → fastest, most frequent during development.
- **Integration tests** → less frequent, usually run in CI/CD pipelines.
- **HTTP testing** → ideal for testing microservices.
- **Benchmarking** → optional but powerful for high-performance Go services.
