# poseidon-cross-platform

[Poseidon](ttps://eprint.iacr.org/2019/458.pdf) is a Zero-Knowledge-Proof friendly hash function widely used in ZK 
applications. Compared with other hash function commonly used in crypto space, Poseidon will incur less circuit 
constraints while preserving security presumptions.

This repository provides examples to access [iden3](https://github.com/iden3/go-iden3-crypto) poseidon implementation 
over BN254 in other platforms through C shared libraries, following [cgo](https://github.com/vladimirvivien/go-cshared-examples)
examples.

## Prerequisites and Dependencies
- Golang 1.20

## Build

1. Clone this repository to your local machine.
2. `cd src`
3. `sh build-so.sh`
4. You can find the dynamic library `poseidon.so` and header file `poseidon.h` under `output/` folder.
5. Alternatively, prebuilt binaries for common platforms are available in the [binaries/](./binaries/) folder of the 
repository.

## C Usage

Check out example to use the shared library under [examples/c](examples/c/)

- `sh build-and-run.sh`
- By default, it is linked to binary under [binaries/macos-arm64/](binaries/macos-arm64/)
- As indicated by the example, C program is responsible to manage the memory of input arguments and output result.




