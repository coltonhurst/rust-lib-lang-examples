# rust-lib-lang-examples

This repository shows a basic example of how to build a Rust library and statically link it with other languages via C FFI. You'll need to build the Rust library before attempting to build the binary programs.

Note: this has been written for Mac & Linux, not Windows.

## Projects

Main Library

- [rustlib](./rustlib): this is the Rust library containing the main implementation
    - Build this with: `cargo build`

Binary Programs

- [cbin](./cbin/): this is the C binary project that uses the rustlib project
    - To build this with gcc: `gcc main.c -L../rustlib/target/debug -lrustlib -o main`
    - To build this with musl on Linux: `musl-gcc main.c -L../rustlib/target/debug -lrustlib -o main -static -Wl,-unresolved-symbols=ignore-all`
- [gobin](./gobin): this is the Go binary project that uses the rustlib project
    - Build this with: `go build`

## Other

### Leaking Paths

While attempting to keep the examples as basic as possible, it might be important to note that some paths are leaked in the built C binary. You can see this by using a binay inspection tool like `fq`.

To see this on a mac:
1. `brew install wader/tap/fq`
2. After building `cbin/main.c`, run: `fq dd ./main | less`

**The Fix**

You can fix this by stripping extra info during the linking process with `-s`, like: `gcc main.c -L../rustlib/target/debug -lrustlib -o main -s`

## Thanks

Shout-out to [tangowithfoxtrot](https://github.com/tangowithfoxtrot) for the pair programming!
