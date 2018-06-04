//build +linux

package srcs

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func CpuModel() ([]byte, error) {
	// TODO: Check for non-procfs FS at /proc and return ErrUnreliableInfo.

	vendorId := ""
	family := ""
	model := ""

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), "vendor_id") {
			vendorId = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
		if strings.HasPrefix(scnr.Text(), "cpu_family") {
			family = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
		if strings.HasPrefix(scnr.Text(), "model") {
			model = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
	}
	if err := scnr.Err(); err != nil && err != io.EOF {
		return nil, err
	}

	return []byte(vendorId + ":" + family + ":" + model), err
}
