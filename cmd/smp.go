// Copyright (c) The kotama Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build sys_emu

package cmd

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/soc/aifoundry/etsoc1/minion"
)

var trail [64]byte

func init() {
	shell.Add(shell.Cmd{
		Name:    "smp",
		Args:    2,
		Pattern: regexp.MustCompile(`^smp (\d+) (\d+)$`),
		Syntax:  "<workloads> <minions>",

		Help: "launch SMP test",
		Fn:   smpCmd,
	})
}

//go:nosplit
func workload() {
	id := RV64.ID()
	trail[id] = 0x21 + byte(id)
}

func smpCmd(console *shell.Interface, arg []string) (string, error) {
	var res bytes.Buffer
	var wg sync.WaitGroup
	var cc sync.Map
	var total int

	n, err := strconv.Atoi(arg[0])

	if err != nil {
		return "", fmt.Errorf("invalid workload count: %v", err)
	}

	ncpu, err := strconv.Atoi(arg[1])

	if err != nil {
		return "", fmt.Errorf("invalid minions count: %v", err)
	}

	if count := minion.NumCPU(); ncpu > count {
		return "", fmt.Errorf("invalid minions count: %d > %d", ncpu, count)
	}

	fmt.Fprintf(console.Output, "Trail %s\n", trail)
	fmt.Fprintf(console.Output, "launching %d workloads on %d minions\n", n, ncpu)

	start := time.Now()

	for i := 0; i < n; i++ {
		wg.Go(func() {
			hart := i % ncpu
			minion.TaskWorkload(hart, workload, true)

			for {
				if actual, loaded := cc.LoadOrStore(hart, 1); loaded {
					if cc.CompareAndSwap(hart, actual.(int), actual.(int)+1) {
						break
					}
				} else {
					break
				}
			}
		})
	}

	wg.Wait()
	elapsed := time.Since(start)

	cc.Range(func(hart any, count any) bool {
		total += count.(int)
		fmt.Fprintf(&res, "hart %2d %2d:%s\n", hart.(int), count.(int), strings.Repeat("░", count.(int)))
		return true
	})

	fmt.Fprintf(&res, "Total  %3d (%v)\n", total, elapsed)
	fmt.Fprintf(&res, "Trail  %s\n", trail)

	for i := range trail {
		trail[i] = 0
	}

	return res.String(), nil
}
