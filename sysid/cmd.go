package main

import (
	"crypto"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/foxcpp/go-sysid"
)

type dummyHash struct {
	res []byte
}

func (h *dummyHash) Reset() {
	h.res = nil
}

func (h *dummyHash) Sum(b []byte) []byte {
	return append(b, h.res...)
}

func (h *dummyHash) Size() int {
	return len(h.res)
}

func (h *dummyHash) BlockSize() int {
	return 0
}

func (h *dummyHash) Write(in []byte) (int, error) {
	h.res = append(h.res, in...)
	return len(in), nil
}

var hashes = map[string]crypto.Hash{
	"md5":         crypto.MD5,
	"sha-1":       crypto.SHA1,
	"sha-224":     crypto.SHA224,
	"sha-256":     crypto.SHA256,
	"sha-384":     crypto.SHA384,
	"sha-512":     crypto.SHA512,
	"sha3-256":    crypto.SHA3_256,
	"sha3-384":    crypto.SHA3_384,
	"sha3-512":    crypto.SHA3_512,
	"blake2b-256": crypto.BLAKE2b_256,
	"blake2b-384": crypto.BLAKE2b_384,
	"blake2b-512": crypto.BLAKE2b_512,
}

func main() {
	rawInfo := flag.Bool("r", false, "print raw information instead of hex-encoded hash")
	hashName := flag.String("a", "sha-512", "use specified hash algorithm")
	flag.Parse()

	hash, prs := hashes[*hashName]
	if !prs {
		fmt.Println("Unknown hash function.")
		fmt.Print("Supported hashes: ")
		for k, _ := range hashes {
			fmt.Print(k, " ")
		}
		os.Exit(2)
	}
	if !hash.Available() {
		fmt.Println("Hash is not availiable in this build.")
		os.Exit(2)
	}

	if *rawInfo {
		id, err := sysid.SysIDCustom(false, &dummyHash{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(id))
	} else {
		id, err := sysid.HWIDCustom(false, hash.New())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(hex.EncodeToString(id))
	}

}
