// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build sifive_u

package cmd

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strconv"
	_ "unsafe"

	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/board/qemu/sifive_u"
	"github.com/usbarmory/tamago/soc/sifive/fu540"
)

//go:linkname ramSize runtime/goos.RamSize
var ramSize uint = 8 << 20

var RV64 = fu540.RV64

func init() {
	Terminal = sifive_u.UART0

	shell.Add(shell.Cmd{
		Name:    "msip",
		Args:    1,
		Pattern: regexp.MustCompile(`^msip (\d+)$`),
		Syntax:  "<hart>",
		Help:    "machine-level software interrupt",
		Fn:      ipiCmd,
	})
}

func infoCmd(_ *shell.Interface, _ []string) (string, error) {
	var res bytes.Buffer

	ramStart, ramEnd := runtime.MemRegion()
	txtStart, txtEnd := runtime.TextRegion()
	datStart, datEnd := runtime.DataRegion()

	name, freq := Target()
	features := fu540.RV64.Features()
	id := fu540.RV64.ID()

	fmt.Fprintf(&res, "SoC ..........: %s @ %v MHz (rv64%s)\n", name, freq/1e6, features.Extensions)
	fmt.Fprintf(&res, "Runtime ......: %s %s/%s thread %d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, id)
	fmt.Fprintf(&res, "RAM ..........: %#08x-%#08x (%d MiB)\n", ramStart, ramEnd, (ramEnd-ramStart)/(1024*1024))
	fmt.Fprintf(&res, "Text .........: %#08x-%#08x (%d KiB)\n", txtStart, txtEnd, (txtEnd-txtStart)/(1024))
	fmt.Fprintf(&res, "Data .........: %#08x-%#08x (%d KiB)\n", datStart, datEnd, (datEnd-datStart)/(1024))
	fmt.Fprintf(&res, "Frequency ....: %v MHz\n", freq/1e6)

	return res.String(), nil
}

func isr() {
	hart := fu540.RV64.ID()
	defer fu540.CLINT.ClearIPI(int(hart))

	log.Printf("got IRQ on hart %d\n", hart)
}

func ipiCmd(_ *shell.Interface, arg []string) (string, error) {
	id, err := strconv.Atoi(arg[0])

	if err != nil {
		return "", fmt.Errorf("invalid Hart ID, %v", err)
	}

	fu540.CLINT.IPI(id)

	return "", nil
}

func Target() (name string, freq uint32) {
	return fu540.Model(), fu540.Freq()
}
