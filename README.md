# Aardvark Prototype Artifact

## Setup

These instructions should work on Ubuntu 18.04.4.

To replicate benchmarks, clone this repository and its submodules.

Build the vector commitment library:
```
$ cd veccom-rust
$ cargo build --release
$ go test -v
```
The last test command should pass if setup succeeded.

Afterwards, go to `go-algorand` and follow the installation instructions there.

### go-algorand

You will need to build libsodium.  The most straightforward way is to install `autoconf`, `automake`, and `libtool` for your system and then run `cd go-algorand; make`.

## Benchmarks

To generate a workload on the machine, run `go-algorand/ledger/bench.sh` (from its directory). Note that this might take a while.

After a workload is generated, you can run benchmarks.

To run validator benchmarks with varying numbers of cores (you'll need `numactl`), run `go-algorand/ledger/cores.sh` (from its directory).

To run the archive benchmark, execute `go-algorand/ledger/acores.sh` (from its directory).
