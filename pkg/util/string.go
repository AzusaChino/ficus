package util

import "unsafe"

// GetString from byte array
func GetString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
