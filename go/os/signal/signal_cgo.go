// +build cgo

package signal

import "C"
import (
	"os"

	"github.com/searKing/golang/go/os/signal/cgo"
)

// signalAction act as signal.Notify, which invokes the Go signal handler.
// https://godoc.org/os/signal#hdr-Go_programs_that_use_cgo_or_SWIG
func signalAction(sigs ...os.Signal) {
	for _, sig := range sigs {
		cgo.SignalAction(Signum(sig))
	}
}

// dumpSignalTo redirects log to fd, stdout if not set.
func dumpSignalTo(fd int) {
	cgo.SetSignalDumpToFd(fd)
}

// dumpStacktraceTo set dump file path of stacktrace when signal is triggered, nop if not set.
func dumpStacktraceTo(name string) {
	cgo.SetBacktraceDumpToFile(name)
}

// dumpPreviousStacktrace dumps human readable stacktrace to fd, which is set by SetSignalDumpToFd.
func dumpPreviousStacktrace() {
	cgo.DumpPreviousStacktrace()
}

// previousStacktrace returns a human readable stacktrace
func previousStacktrace() string {
	return cgo.PreviousStacktrace()
}