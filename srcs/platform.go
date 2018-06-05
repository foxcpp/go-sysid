// +build linux

package srcs

func Platform() ([]byte, error) {
	return []byte("linux"), nil
}
