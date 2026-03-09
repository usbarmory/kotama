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

Building
--------

This project requires the experimental `GOOS=tamago` compiler branch
[tamago1.26.0-73608-softfloat branch](https://github.com/abarisani/tamago-go/tree/tamago1.26.1-73608-softfloat).

You can build, and run in QEMU, as follows:

```
./run.sh
```

Operation
=========

```
build                                    # build information
date                                     # show runtime date and time
exit                                     # close session
halt                                     # halt the machine
help                                     # this help
info                                     # device information
kem                                      # benchmark post-quantum KEM
ls              (<path>)?                # list directory contents
metrics                                  # show runtime metrics
peek            <hex addr> <size>        # memory display (use with caution)
poke            <hex addr> <hex value>   # memory write   (use with caution)
rand                                     # gather 32 random bytes
reboot                                   # reset device
stack                                    # goroutine stack trace (current)
stackall                                 # goroutine stack trace (all)
test                                     # launch tests
uptime                                   # show system running time

> info
CPU ..........: rv64cfimsu
Runtime ......: go1.26.1 tamago/riscv64 GOMAXPROCS=1
RAM ..........: 0x80000000-0x80600000 (6 MiB)
Text .........: 0x80010000-0x800f1de8 (903 KiB)
Data .........: 0x8023e300-0x802729e0 (209 KiB)
```

License
=======

Copyright (c) The kotama authors. All Rights Reserved.

These source files are distributed under the BSD-style license found in the
[LICENSE](https://github.com/usbarmory/kotama/blob/main/LICENSE) file.
