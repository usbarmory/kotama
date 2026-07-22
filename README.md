Introduction
============

```
こたま | kotama | noun

Small-sized egg; small sphere.
An egg that is smaller than the standard grade.
```

This is an experiment for bare metal Go unikernels, as supported by
[TamaGo](https://github.com/usbarmory/tamago) on RISCV64 low-memory, extension
constrained targets.

This project requires the experimental `GOOS=tamago` compiler branch
[tamago1.26.5-softfloat branch](https://github.com/abarisani/tamago-go/tree/tamago1.26.5-softfloat),
such branch extends mainline [tamago](https://github.com/usbarmory/tamago/wiki)
to support the following:

  * `GOSOFT=1`: compiler build time variable to enable soft float for `riscv64`, removing
    requirement for `ad` extensions and forcing single-threaded operation.

  * `tiny`: tamago library build tag for considerable reduction of RAM allocation requirements.

Supported RISC-V targets
========================

* [AI Foundry Erbium](https://github.com/aifoundry-org/erbium)
* [AI Foundry ET-SoC-1](https://github.com/aifoundry-org/et-man)
* [SiFive FU540](https://www.qemu.org/docs/master/system/riscv/sifive_u.html)

Building the compiler
---------------------

```
wget https://github.com/abarisani/tamago-go/archive/refs/heads/tamago1.26.5-softfloat.zip
unzip tamago1.26.5-softfloat.zip
cd tamago-go-tamago1.26.5-softfloat/src && ./all.bash
cd ../bin && export TAMAGO=`pwd`/go
```

> [!NOTE]
> As `GOSOFT=1` not accounted in Go build ID, it is advised to
> `go clean -cache` before compiling for the first time with the `GOSOFT=1`
> enabled compiler.

Building and running
--------------------

You can build, and run the available targets as follows:

```
# AI Foundry Erbium processor (16MB RAM) with /opt/et/bin/erbium_emu
./run erbium

# AI Foundry ET-SoC-1 Minion Core (16MB RAM) with /opt/et/bin/sys_emu
./run etsoc1

# SiFive FU540 (6MB RAM) with qemu-system-riscv64
./run fu540
```

> [!TIP]
> Extra arguments (e.g. `-gdb` for ET emulators) can be passed and are applied
> to the emulator command line.

Operation
=========

```
build                                    # build information
cat         <path>                       # show file contents
date        (<time in RFC339 format>)?   # show/change runtime date and time
exit                                     # close session
halt                                     # halt the machine
help                                     # this help
info                                     # device information
kem                                      # benchmark post-quantum KEM
ls          (<path>)?                    # list directory contents
metrics                                  # show runtime metrics
msip        <hart>                       # machine-level software interrupt
peek        <hex addr> <size>            # memory display (use with caution)
poke        <hex addr> <hex value>       # memory write   (use with caution)
rand                                     # gather 32 random bytes
reset       (soft|warm)?                 # reset system
smp         <workloads> <minions>        # launch SMP test
stack                                    # goroutine stack trace (current)
stackall                                 # goroutine stack trace (all)
test                                     # launch tests
uptime                                   # show system running time

> info
SoC ..........: ET-SoC-1 Minion Core @ 200 MHz (rv64cfimsux)
Minions ......: 64
Runtime ......: go1.26.5 tamago/riscv64 thread 0
RAM ..........: 0x8000000000-0x8400000000 (16384 MiB)
Text .........: 0x8000010000-0x80000f8be8 (930 KiB)
Data .........: 0x800024e340-0x8000283588 (212 KiB)

> smp 100 4
launching 100 workloads on 4 minions
hart  0 25:░░░░░░░░░░░░░░░░░░░░░░░░░
hart  3 25:░░░░░░░░░░░░░░░░░░░░░░░░░
hart  1 25:░░░░░░░░░░░░░░░░░░░░░░░░░
hart  2 25:░░░░░░░░░░░░░░░░░░░░░░░░░
Total  100 (474.35ms)
```

License
=======

Copyright (c) The kotama authors. All Rights Reserved.

These source files are distributed under the BSD-style license found in the
[LICENSE](https://github.com/usbarmory/kotama/blob/main/LICENSE) file.
