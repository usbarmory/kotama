// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"time"
)

const sleep = 100 * time.Millisecond

func timerTest() (tag string, res string) {
	start := time.Now()
	t := time.NewTimer(sleep)

	for now := range t.C {
		tag = fmt.Sprintf("timer event %v (actual %v)", sleep, now.Sub(start))
		break
	}

	return
}

func sleepTest() (tag string, res string) {
	start := time.Now()
	time.Sleep(sleep)

	tag = fmt.Sprintf("timer sleep %s (actual %v)", sleep, time.Since(start))

	return
}
