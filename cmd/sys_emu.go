// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build sys_emu

package cmd

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/board/aifoundry/sys_emu"
	"github.com/usbarmory/tamago/soc/aifoundry/etsoc1/minion"
)

func init() {
	Terminal = sys_emu.UART0
}

func date(epoch int64) {
	minion.RV64.SetTime(epoch)
}

func uptime() (ns int64) {
	return minion.RV64.GetTime() - minion.RV64.TimerOffset
}

func infoCmd(_ *shell.Interface, _ []string) (string, error) {
	var res bytes.Buffer

	ramStart, ramEnd := runtime.MemRegion()
	txtStart, txtEnd := runtime.TextRegion()
	datStart, datEnd := runtime.DataRegion()

	name, freq := Target()
	features :=  minion.RV64.Features()

	fmt.Fprintf(&res, "SoC ..........: %s @ %v MHz (rv64%s)\n", name, freq/1e6, features.Extensions)
	fmt.Fprintf(&res, "Runtime ......: %s %s/%s thread %d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, features.HartID)
	fmt.Fprintf(&res, "RAM ..........: %#08x-%#08x (%d MiB)\n", ramStart, ramEnd, (ramEnd-ramStart)/(1024*1024))
	fmt.Fprintf(&res, "Text .........: %#08x-%#08x (%d KiB)\n", txtStart, txtEnd, (txtEnd-txtStart)/(1024))
	fmt.Fprintf(&res, "Data .........: %#08x-%#08x (%d KiB)\n", datStart, datEnd, (datEnd-datStart)/(1024))

	return res.String(), nil
}

func Target() (name string, freq uint32) {
	name = minion.Model()
	freq = minion.CoreFreq

	return
}
