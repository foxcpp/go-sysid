// +build linux

package srcs

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func OSFlavor() ([]byte, error) {
	osrel, err := os.Open("/etc/os-release")
	if err != nil {
		return []byte("UNKNOWN_OS_FLAVOR"), nil
	}

	scnr := bufio.NewScanner(osrel)
	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), "NAME=\"") {
			return scnr.Bytes()[6 : len(scnr.Bytes())-1], nil
		}
	}
	if err := scnr.Err(); err != nil && err != io.EOF {
		return nil, err
	}
	return []byte("UNKNOWN_OS_FLAVOR"), nil
}
