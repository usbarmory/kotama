set -x
GOOS=tamago GOARCH=riscv64 GOSOFT=1 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -tags erbium_emu,tiny,semihosting -trimpath -ldflags "-T 0x40010000 -R 0x1000" main.go && \
RT0=$(riscv64-linux-gnu-readelf -a main|grep -i 'Entry point' | cut -dx -f2) && \
/opt/et/bin/erbium_emu -elf_load /mnt/git/public/kotama/main -reset_pc $RT0 -max_cycles -1 -single_thread
