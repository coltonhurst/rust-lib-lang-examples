# rust-lib-lang-examples

This repository shows a basic example of how to build a Rust library and statically link it with other languages via C FFI.

## Projects

Main Library

- [rustlib](./rustlib): this is the Rust library containing the main implementation
    - Build this with: `cargo build`

Binary Programs

- [cbin](./cbin/): this is the C binary project that uses the rustlib project
    - Build this with: `gcc main.c -L../rustlib/target/debug -lrustlib -o main`
- [gobin](./gobin): this is the Go binary project that uses the rustlib project
    - Build this with: `go build`
