package utils

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGID 获取当前goroutine的ID,性能一般,节制调用
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
