// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"runtime/debug"
	"runtime/goos"
	"runtime/pprof"
	"time"

	"github.com/usbarmory/tamago-example/shell"
)

const maxBufferSize = 102400

var Terminal io.ReadWriter

func init() {
	shell.Add(shell.Cmd{
		Name: "build",
		Help: "build information",
		Fn:   buildInfoCmd,
	})

	shell.Add(shell.Cmd{
		Name: "exit",
		Help: "close session",
		Fn:   exitCmd,
	})

	shell.Add(shell.Cmd{
		Name: "halt",
		Help: "halt the machine",
		Fn:   haltCmd,
	})

	shell.Add(shell.Cmd{
		Name: "stack",
		Help: "goroutine stack trace (current)",
		Fn:   stackCmd,
	})

	shell.Add(shell.Cmd{
		Name: "stackall",
		Help: "goroutine stack trace (all)",
		Fn:   stackallCmd,
	})

	shell.Add(shell.Cmd{
		Name: "date",
		Args: 1,
		Help: "show runtime date and time",
		Fn:   dateCmd,
	})

	shell.Add(shell.Cmd{
		Name: "uptime",
		Help: "show system running time",
		Fn:   uptimeCmd,
	})

	// The following commands are board specific, therefore their Fn
	// pointers are defined elsewhere in the respective target files.

	shell.Add(shell.Cmd{
		Name: "info",
		Help: "device information",
		Fn:   infoCmd,
	})

	shell.Add(shell.Cmd{
		Name: "reboot",
		Help: "reset device",
		Fn:   rebootCmd,
	})
}

func buildInfoCmd(_ *shell.Interface, _ []string) (string, error) {
	var res bytes.Buffer

	if bi, ok := debug.ReadBuildInfo(); ok {
		res.WriteString(bi.String())
	}

	return res.String(), nil
}

func exitCmd(console *shell.Interface, _ []string) (string, error) {
	fmt.Fprintf(console.Output, "Goodbye from %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return "logout", io.EOF
}

func haltCmd(console *shell.Interface, _ []string) (string, error) {
	fmt.Fprintf(console.Output, "Goodbye from %s/%s\n", runtime.GOOS, runtime.GOARCH)

	time.AfterFunc(
		100*time.Millisecond,
		func() { goos.Exit(0) },
	)

	return "halted", io.EOF
}

func stackCmd(_ *shell.Interface, _ []string) (string, error) {
	return string(debug.Stack()), nil
}

func stackallCmd(_ *shell.Interface, _ []string) (string, error) {
	buf := new(bytes.Buffer)
	pprof.Lookup("goroutine").WriteTo(buf, 1)

	return buf.String(), nil
}

func dateCmd(_ *shell.Interface, arg []string) (res string, err error) {
	return fmt.Sprintf("%s", time.Now().Format(time.RFC3339)), nil
}

func uptimeCmd(_ *shell.Interface, _ []string) (string, error) {
	return fmt.Sprintf("%s", time.Duration(uptime())*time.Nanosecond), nil
}
