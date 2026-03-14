set -x
GOOS=tamago GOARCH=riscv64 GOSOFT=1 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -tags sifive_u,tiny,semihosting,linkramsize -trimpath -ldflags "-T 0x80010000 -R 0x1000" main.go && \
RT0=$(nm main|grep _rt0_riscv64_tamago | cut -d' ' -f1) && \
echo ".equ RT0_RISCV64_TAMAGO, 0x$RT0" > ${PWD}/tools/bios.cfg && \
cd ${PWD}/tools && ./build_riscv64_bios.sh && \
cd ../ && \
# rv64imfc
qemu-system-riscv64 \
  -machine sifive_u -cpu rv64,a=off,d=off,h=off,s=on,u=on,zawrs=off -m 6M \
  -nographic -monitor none -semihosting -serial stdio -net none \
  -dtb ${PWD}/tools/qemu.dtb -bios ${PWD}/tools/bios.bin -kernel main
