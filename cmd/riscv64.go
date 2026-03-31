// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build riscv64

package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	_ "unsafe"

	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/riscv64"
)

func init() {
	go RV64.ServiceInterrupts(isr)
	RV64.EnableInterrupts()

	shell.Add(shell.Cmd{
		Name:    "msip",
		Args:    1,
		Pattern: regexp.MustCompile(`^msip (\d+)$`),
		Syntax:  "<hart>",
		Help:    "machine-level software interrupt",
		Fn:      ipiCmd,
	})
}

func date(epoch int64) {
	RV64.SetTime(epoch)
}

func uptime() (ns int64) {
	return RV64.GetTime() - RV64.TimerOffset
}

func isr() {
	hart := RV64.ID()
	defer riscv64.ClearIPI(int(hart))

	log.Printf("got IRQ on hart %d\n", hart)
}

func ipiCmd(_ *shell.Interface, arg []string) (string, error) {
	id, err := strconv.Atoi(arg[0])

	if err != nil {
		return "", fmt.Errorf("invalid Hart ID, %v", err)
	}

	riscv64.IPI(id)

	return "", nil
}
