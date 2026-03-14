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
[tamago1.26.1-73608-softfloat branch](https://github.com/abarisani/tamago-go/tree/tamago1.26.1-73608-softfloat),
such branch extends mainline [tamago](https://github.com/usbarmory/tamago/wiki)
to support the following:

  * `GOSOFT=1`: compiler build time variable to enable soft float for `riscv64`, removing
    requirement for `ad` extensions and forcing single-threaded operation.

  * `tiny`: tamago library build tag for considerable reduction of RAM allocation requirements.

Supported RISC-V targets
========================

* [AI Foundry Erbium](https://github.com/aifoundry-org/erbium)
* [SiFive FU540](https://www.qemu.org/docs/master/system/riscv/sifive_u.html)

Building the compiler
---------------------

```
wget https://github.com/abarisani/tamago-go/archive/refs/heads/tamago1.26.1-73608-softfloat.zip
unzip tamago1.26.1-73608-softfloat.zip
cd tamago-go-tamago1.26.1-73608-softfloat/src && ./all.bash
cd ../bin && export TAMAGO=`pwd`/go
```

Building the application
------------------------

You can build, and run for two separate emulated targets, as follows:

```
# AI Foundry Erbium processor
./run-erbium_emu.sh

# SiFive FU540
./run-sifive_u.sh
```

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
peek        <hex addr> <size>            # memory display (use with caution)
poke        <hex addr> <hex value>       # memory write   (use with caution)
rand                                     # gather 32 random bytes
reboot                                   # reset device
stack                                    # goroutine stack trace (current)
stackall                                 # goroutine stack trace (all)
test                                     # launch tests
uptime                                   # show system running time

> info
SoC ..........: FU540 @ 999 MHz (rv64cfimsu)
Runtime ......: go1.26.1 tamago/riscv64 GOMAXPROCS=1
RAM ..........: 0x80000000-0x80600000 (6 MiB)
Text .........: 0x80010000-0x800f5898 (918 KiB)
Data .........: 0x80245300-0x80279c20 (210 KiB)
Frequency ....: 999 MHz
```

License
=======

Copyright (c) The kotama authors. All Rights Reserved.

These source files are distributed under the BSD-style license found in the
[LICENSE](https://github.com/usbarmory/kotama/blob/main/LICENSE) file.
