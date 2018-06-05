// +build linux

package srcs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Sadly, amount of information about memory available in Linux to non-root
// users is limited to memory size. We could get more if running as setuid
// binary. (TODO!)

func MemoryInfo() ([]byte, error) {
	meminfo, err := os.Open("/proc/meminfo")
	if err != nil {
		return []byte("NO_MEM_INFO"), nil
	}
	scnr := bufio.NewScanner(meminfo)

	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), "MemTotal") {
			return bytes.TrimSpace(bytes.Split(scnr.Bytes(), []byte(":"))[1]), nil
		}
	}
	if err := scnr.Err(); err != nil && err != io.EOF {
		return nil, fmt.Errorf("MemoryInfo: %v", err)
	}
	return nil, errors.New("MemoryInfo: no MemTotal found in /proc/meminfo")
}
