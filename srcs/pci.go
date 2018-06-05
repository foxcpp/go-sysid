// +build linux

package srcs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

const pciBaseDir = "/sys/bus/pci/devices"

func PCIDeviceList() ([]byte, error) {
	// TODO: Check for sysfs mount.

	files, err := ioutil.ReadDir(pciBaseDir)
	if err != nil {
		return []byte("NO_PCI_INFO"), nil
	}

	sort.Slice(files, func(a, b int) bool {
		return files[a].Name() < files[b].Name()
	})

	ids := [][]byte{}

	for _, file := range files {
		var err error
		class, vendor, device := []byte{}, []byte{}, []byte{}

		class, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "class"))
		if err != nil {
			return nil, fmt.Errorf("PCIDeviceList: %v", err)
		}
		class = bytes.TrimSpace(class)
		vendor, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "vendor"))
		if err != nil {
			return nil, fmt.Errorf("PCIDeviceList: %v", err)
		}
		vendor = bytes.TrimSpace(vendor)
		device, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "device"))
		if err != nil {
			return nil, fmt.Errorf("PCIDeviceList: %v", err)
		}
		device = bytes.TrimSpace(device)

		ids = append(ids, bytes.Join([][]byte{class, vendor, device}, []byte(":")))
	}

	return bytes.Join(ids, []byte(":")), nil
}
