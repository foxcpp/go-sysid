// +build linux

package srcs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/foxcpp/go-sysid/mountcheck"
)

func CpuModel() ([]byte, error) {
	var unreliable error
	if !mountcheck.IsProc("/proc") {
		unreliable = ErrUnreliableInfo
	}

	vendorId := ""
	family := ""
	model := ""

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return []byte("NO_CPU_INFO"), unreliable
	}

	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), "vendor_id") {
			vendorId = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
		if strings.HasPrefix(scnr.Text(), "cpu family") {
			family = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
		if strings.HasPrefix(scnr.Text(), "model\t\t") { // tabs to not touch "model name".
			model = strings.TrimSpace(strings.Split(scnr.Text(), ":")[1])
		}
	}
	if err := scnr.Err(); err != nil && err != io.EOF {
		return nil, fmt.Errorf("CpuModel: %v", err)
	}

	return []byte(vendorId + ":" + family + ":" + model), unreliable
}
