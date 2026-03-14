// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/usbarmory/kotama/cmd"

	"github.com/usbarmory/tamago-example/shell"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	_, ramEnd := runtime.MemRegion()
	dataStart, _ := runtime.DataRegion()
	memoryLimit := float64(ramEnd-dataStart) * 0.90
	debug.SetMemoryLimit(int64(math.Round(memoryLimit)))
}

func main() {
	banner := fmt.Sprintf("%s/%s (%s) • %s",
		runtime.GOOS, runtime.GOARCH, runtime.Version(), "こたま")

	console := &shell.Interface{
		Banner:     banner,
		ReadWriter: cmd.Terminal,
	}

	console.Start(cmd.IsVT100)

	log.Printf("exit")
}
