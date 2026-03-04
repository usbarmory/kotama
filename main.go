package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/debug"
	_ "unsafe"

	_ "github.com/usbarmory/kotama/cmd"
	"github.com/usbarmory/tamago-example/shell"

	"github.com/usbarmory/tamago/board/qemu/sifive_u"
)

//go:linkname ramSize runtime/goos.RamSize
var ramSize uint = 6 << 20

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
		ReadWriter: sifive_u.UART0,
	}

	console.Start(true)

	log.Printf("exit")
}
