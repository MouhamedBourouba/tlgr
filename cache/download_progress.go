package cache

import (
	"fmt"
	"strings"
)

type downloadProgressCounter struct {
	TotalBytes      uint64
	TransferedBytes uint64
}

func newDownloadProgressCounter(totalBytes uint64) downloadProgressCounter {
	return downloadProgressCounter{TotalBytes: totalBytes, TransferedBytes: 0}
}

func (w *downloadProgressCounter) Write(p []byte) (n int, err error) {
	w.TransferedBytes = w.TransferedBytes + uint64(len(p))
	w.printProgress()
	if w.TransferedBytes >= w.TotalBytes {
		println()
	}
	return len(p), nil
}

func (w downloadProgressCounter) printProgress() {
	fmt.Print("\r", strings.Repeat(" ", 120))
	fmt.Print("\rDownloading Repository %", (float32(w.TransferedBytes)/float32(w.TotalBytes))*100)
}
