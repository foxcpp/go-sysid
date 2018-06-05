package mountcheck

import "os"
import "syscall"

func IsProc(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.Sys().(*syscall.Stat_t).Ino == 1 && fi.Sys().(*syscall.Stat_t).Dev == 4
}
