package srcs

import (
	"io/ioutil"
	"path/filepath"
	"sort"
)

const pciBaseDir = "/sys/bus/pci/devices"

func PCIDeviceList() ([]byte, error) {
	// TODO: Check for sysfs mount.

	files, err := ioutil.ReadDir(pciBaseDir)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(a, b int) bool {
		return files[a].Name() < files[b].Name()
	})

	result := []byte{}

	for _, file := range files {
		var err error
		class, vendor, device := []byte{}, []byte{}, []byte{}

		class, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "class"))
		if err != nil {
			return nil, err
		}
		vendor, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "vendor"))
		if err != nil {
			return nil, err
		}
		device, err = ioutil.ReadFile(filepath.Join(pciBaseDir, file.Name(), "device"))
		if err != nil {
			return nil, err
		}

		result = append(result, class...)
		result = append(result, byte(':'))
		result = append(result, vendor...)
		result = append(result, byte(':'))
		result = append(result, device...)
		result = append(result, byte(':'))
	}

	return result, nil
}
