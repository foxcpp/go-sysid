package sysid

import (
	"crypto/sha512"
	"hash"

	"github.com/foxcpp/go-sysid/srcs"
)

var ErrUnreliableInfo = srcs.ErrUnreliableInfo

var sources = []func() ([]byte, error){
	srcs.Platform,
	srcs.OSFlavor,
	srcs.CpuModel,
	srcs.PCIDeviceList,
	srcs.MemoryInfo,
	srcs.ATADeviceNames,
	srcs.MACAddresses,
}

/*
Collect system information and apply custom hash h on it.

allowUnreliable is currently not used.
*/
func SysIDCustom(allowUnreliable bool, h hash.Hash) ([]byte, error) {
	for _, src := range sources {
		res, err := src()
		if err != nil {
			if !(allowUnreliable && err == ErrUnreliableInfo) {
				return nil, err
			}
		}
		res = append(res, byte(':'))
		if _, err := h.Write(res); err != nil {
			return nil, err
		}
	}
	return h.Sum([]byte{}), nil
}

/*
Collect system information and apply SHA-512 on it.
*/
func SysID() ([]byte, error) {
	return SysIDCustom(false, sha512.New())
}
