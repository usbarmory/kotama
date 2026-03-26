// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build erbium_emu

package cmd

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"strconv"

	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/board/aifoundry/erbium_emu"
	"github.com/usbarmory/tamago/soc/aifoundry/erbium"
)

var RV64 = erbium.RV64

func init() {
	Terminal = erbium_emu.UART0

	shell.Add(shell.Cmd{
		Name:    "msip",
		Args:    1,
		Pattern: regexp.MustCompile(`^msip (\d+)$`),
		Syntax:  "<hart>",
		Help:    "machine-level software interrupt",
		Fn:      ipiCmd,
	})

	shell.Add(shell.Cmd{
		Name:    "reset",
		Args:    1,
		Pattern: regexp.MustCompile(`^reset(?: (cold|soft))?$`),
		Help:    "reset system",
		Syntax:  "(soft|warm)?",
		Fn:      resetCmd,
	})
}

func infoCmd(_ *shell.Interface, _ []string) (string, error) {
	var res bytes.Buffer

	ramStart, ramEnd := runtime.MemRegion()
	txtStart, txtEnd := runtime.TextRegion()
	datStart, datEnd := runtime.DataRegion()

	name, version, freq := Target()
	features :=  erbium.RV64.Features()
	id := erbium.RV64.ID()

	fmt.Fprintf(&res, "SoC ..........: %s (%x) @ %v MHz (rv64%s)\n", name, version, freq/1e6, features.Extensions)
	fmt.Fprintf(&res, "Runtime ......: %s %s/%s thread %d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, id)
	fmt.Fprintf(&res, "RAM ..........: %#08x-%#08x (%d MiB)\n", ramStart, ramEnd, (ramEnd-ramStart)/(1024*1024))
	fmt.Fprintf(&res, "Text .........: %#08x-%#08x (%d KiB)\n", txtStart, txtEnd, (txtEnd-txtStart)/(1024))
	fmt.Fprintf(&res, "Data .........: %#08x-%#08x (%d KiB)\n", datStart, datEnd, (datEnd-datStart)/(1024))

	return res.String(), nil
}

func isr() {
	hart := erbium.RV64.ID()
	defer erbium.ClearIPI(int(hart))

	fmt.Printf("got IRQ on hart %d\n", hart)
}

func ipiCmd(_ *shell.Interface, arg []string) (string, error) {
	id, err := strconv.Atoi(arg[0])

	if err != nil {
		return "", fmt.Errorf("invalid Hart ID, %v", err)
	}

	erbium.IPI(id)

	return "", nil
}

func resetCmd(_ *shell.Interface, arg []string) (_ string, err error) {
	erbium.Reset(arg[0] == "soft")
	return
}

func Target() (name string, version uint32, freq uint32) {
	name, version = erbium.Model()
	freq = erbium.CoreFreq
	return
}
