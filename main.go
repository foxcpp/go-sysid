package hwid

import (
	"crypto/sha512"
	"hash"

	"github.com/foxcpp/go-hwid/srcs"
)

var ErrUnreliableInfo = srcs.ErrUnreliableInfo

var sources = []func() ([]byte, error){
	srcs.Platform,
	srcs.CpuModel,
	srcs.PCIDeviceList,
}

func HWIDCustom(allowUnreliable bool, combiner hash.Hash) ([]byte, error) {
	for _, src := range sources {
		res, err := src()
		if err != nil {
			if !(allowUnreliable && err == ErrUnreliableInfo) {
				return nil, err
			}
		}
		if _, err := combiner.Write(res); err != nil {
			return nil, err
		}
	}
	return combiner.Sum([]byte{}), nil
}

func HWID() ([]byte, error) {
	return HWIDCustom(false, sha512.New())
}
