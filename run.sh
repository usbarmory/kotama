target="$1"

shift 1

if [ "$target" == "erbium" ]; then
	set -x
	GOOS=tamago GOARCH=riscv64 GOSOFT=1 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -tags erbium_emu,tiny,linkcpuinit -trimpath -ldflags "-T 0x40010000 -R 0x1000" main.go && \
	RT0=$(nm main|grep _rt0_riscv64_tamago | cut -d' ' -f1) && \
	/opt/et/bin/erbium_emu -elf_load main -reset_pc $RT0 -max_cycles -1 -minions 0xff $@
elif [ "$target" == "etsoc1" ]; then
	set -x
	GOOS=tamago GOARCH=riscv64 GOSOFT=1 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -tags sys_emu,linkcpuinit -trimpath -ldflags "-T 0x8000010000 -R 0x1000" main.go && \
	RT0=$(nm main|grep _rt0_riscv64_tamago | cut -d' ' -f1) && \
	/opt/et/bin/sys_emu -elf_load main -reset_pc $RT0 -max_cycles -1 -minions 0xff $@
elif [ "$target" == "fu540" ]; then
	set -x
	GOOS=tamago GOARCH=riscv64 GOSOFT=1 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -tags sifive_u,tiny,semihosting,linkramsize -trimpath -ldflags "-T 0x80010000 -R 0x1000" main.go && \
	RT0=$(nm main|grep _rt0_riscv64_tamago | cut -d' ' -f1) && \
	echo ".equ RT0_RISCV64_TAMAGO, 0x$RT0" > ${PWD}/tools/bios.cfg && \
	cd ${PWD}/tools && ./build_riscv64_bios.sh sifive_u && \
	cd ../ && \
	# rv64imfc
	qemu-system-riscv64 \
	  -machine sifive_u -cpu rv64,a=off,d=off,h=off,s=on,u=on,zawrs=off -m 8M \
	  -nographic -monitor none -semihosting -serial stdio -net none \
	  -dtb ${PWD}/tools/qemu.dtb -bios ${PWD}/tools/sifive_u.bin -kernel main $@
else 
	echo "invalid target, choose: erbium, etsoc1, fu540"
fi
