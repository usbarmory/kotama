// Copyright (c) The TamaGo Authors. All Rights Reserved.
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

	"github.com/usbarmory/tamago-example/shell"

	_ "github.com/usbarmory/tamago/board/aifoundry/erbium_emu"
	"github.com/usbarmory/tamago/soc/aifoundry/erbium"
)

func init() {
	Terminal = erbium.UART0

	shell.Add(shell.Cmd{
		Name:    "reset",
		Args:    1,
		Pattern: regexp.MustCompile(`^reset(?: (cold|soft))?$`),
		Help:    "reset system",
		Syntax:  "(soft|warm)?",
		Fn:      resetCmd,
	})
}

func date(epoch int64) {
	erbium.RV64.SetTime(epoch)
}

func uptime() (ns int64) {
	return erbium.RV64.GetTime() - erbium.RV64.TimerOffset
}

func infoCmd(_ *shell.Interface, _ []string) (string, error) {
	var res bytes.Buffer

	ramStart, ramEnd := runtime.MemRegion()
	txtStart, txtEnd := runtime.TextRegion()
	datStart, datEnd := runtime.DataRegion()

	name, version, freq := Target()

	fmt.Fprintf(&res, "SoC ..........: %s (%x) @ %v MHz (rv64%s)\n", name, version, freq/1e6, erbium.RV64.Features().Extensions)
	fmt.Fprintf(&res, "Runtime ......: %s %s/%s GOMAXPROCS=%d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(-1))
	fmt.Fprintf(&res, "RAM ..........: %#08x-%#08x (%d MiB)\n", ramStart, ramEnd, (ramEnd-ramStart)/(1024*1024))
	fmt.Fprintf(&res, "Text .........: %#08x-%#08x (%d KiB)\n", txtStart, txtEnd, (txtEnd-txtStart)/(1024))
	fmt.Fprintf(&res, "Data .........: %#08x-%#08x (%d KiB)\n", datStart, datEnd, (datEnd-datStart)/(1024))

	return res.String(), nil
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
