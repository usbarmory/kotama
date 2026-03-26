// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build riscv64

package cmd

import (
	_ "unsafe"
)

func init() {
	go RV64.ServiceInterrupts(isr)
	RV64.EnableInterrupts()
}

func date(epoch int64) {
	RV64.SetTime(epoch)
}

func uptime() (ns int64) {
	return RV64.GetTime() - RV64.TimerOffset
}
