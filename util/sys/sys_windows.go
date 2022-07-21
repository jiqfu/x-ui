//go:build windows
// +build windows

package sys

func GetTCPCount() (int, error) {
	return 1, nil
}

func GetUDPCount() (int, error) {
	return 1, nil
}
