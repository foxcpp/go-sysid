package hwid

import (
	"crypto/sha512"
	"hash"

	"github.com/foxcpp/go-hwid/srcs"
)

var ErrUnreliableInfo = srcs.ErrUnreliableInfo

var sources = []func() ([]byte, error){
	srcs.Platform,
	srcs.OSFlavor,
	srcs.CpuModel,
	srcs.PCIDeviceList,
	srcs.MemoryInfo,
	srcs.ATADeviceNames,
}

func HWIDCustom(allowUnreliable bool, combiner hash.Hash) ([]byte, error) {
	for _, src := range sources {
		res, err := src()
		if err != nil {
			if !(allowUnreliable && err == ErrUnreliableInfo) {
				return nil, err
			}
		}
		res = append(res, byte(':'))
		if _, err := combiner.Write(res); err != nil {
			return nil, err
		}
	}
	return combiner.Sum([]byte{}), nil
}

func HWID() ([]byte, error) {
	return HWIDCustom(false, sha512.New())
}
