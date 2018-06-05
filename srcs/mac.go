package srcs

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"syscall"
	"unsafe"
)

type SockaddrHw struct {
	family uint16
	data   [6]byte
	pad    [8]byte
}

type Ifreq struct {
	ifname [16]byte
	hwaddr SockaddrHw
	pad    [8]byte
}

func interfaces() ([]string, error) {
	// TODO: Test for procfs mount.

	dev, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	scnr := bufio.NewScanner(dev)
	// Skip first two lines, table headers.
	scnr.Scan()
	scnr.Scan()

	res := []string{}
	for scnr.Scan() {
		dev := strings.TrimSpace(strings.Split(scnr.Text(), ":")[0])
		// wlan, enpXsX, ethX. Other devices usually created dynamically and
		// should not be considered.
		if strings.HasPrefix(dev, "wlan") || strings.HasPrefix(dev, "enp") || strings.HasPrefix(dev, "eth") {
			res = append(res, dev)
		}

	}
	if err := scnr.Err(); err != nil && err != io.EOF {
		return nil, err
	}
	return res, nil
}

func macAddress(inter string) ([]byte, error) {
	out := Ifreq{}
	copy(out.ifname[:], []byte(inter))
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.Close(fd)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.SIOCGIFHWADDR), uintptr(unsafe.Pointer(&out)))
	if errno != 0 {
		return nil, errno
	}
	return []byte(hex.EncodeToString(out.hwaddr.data[:])), nil
}

func MACAddresses() ([]byte, error) {
	addresses := [][]byte{}
	interfs, err := interfaces()
	if err != nil {
		return nil, fmt.Errorf("MACAddresses: %v", err)
	}
	for _, interf := range interfs {
		addr, err := macAddress(interf)
		if err != nil {
			return nil, fmt.Errorf("MACAddresses: %v", err)
		}
		addresses = append(addresses, []byte(addr))
	}

	// Sort addreses, so order will not get mixed in some odd configurations.
	sort.Slice(addresses, func(a, b int) bool {
		return string(addresses[a]) < string(addresses[b])
	})

	return bytes.Join(addresses, []byte(":")), nil
}
