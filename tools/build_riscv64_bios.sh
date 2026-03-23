riscv64-linux-gnu-gcc -march=rv64g -mabi=lp64 -static -mcmodel=medany -fvisibility=hidden -nostdlib -nostartfiles -T$1.ld $1.s -o $1.bin
