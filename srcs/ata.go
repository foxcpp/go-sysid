package srcs

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"
)

func ATADeviceNames() ([]byte, error) {
	// TODO: Check for mountpoint.

	files, err := ioutil.ReadDir("/dev/disk/by-id")
	if err != nil {
		return []byte("NO_ATA_INFO"), nil
	}

	sort.Slice(files, func(a, b int) bool {
		return files[a].Name() < files[b].Name()
	})

	deviceNames := [][]byte{}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "ata-") && !strings.Contains(file.Name(), "-part") {
			deviceNames = append(deviceNames, []byte(strings.Join(strings.Split(file.Name(), "-")[1:], "-")))
		}
	}
	return bytes.Join(deviceNames, []byte(":")), nil
}
