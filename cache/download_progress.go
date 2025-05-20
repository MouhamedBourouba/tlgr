package cache

import (
	"fmt"
	"strings"
)

type downloadProgressCounterWriter struct {
	TotalBytes      uint64
	TransferedBytes uint64
}

func newDownloadProgressCounter(totalBytes uint64) downloadProgressCounterWriter {
	return downloadProgressCounterWriter{TotalBytes: totalBytes, TransferedBytes: 0}
}

func (w *downloadProgressCounterWriter) Write(p []byte) (n int, err error) {
	w.TransferedBytes = w.TransferedBytes + uint64(len(p))
	w.printProgress()
	return len(p), nil
}

func (w downloadProgressCounterWriter) printProgress() {
	fmt.Print("\r", strings.Repeat(" ", 120))
	fmt.Print("\rDownloading Repository %", (float32(w.TransferedBytes)/float32(w.TotalBytes))*100)

	if w.TransferedBytes >= w.TotalBytes {
		fmt.Print("\n")
	}
}
