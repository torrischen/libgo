package converter

import "unsafe"

func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{s, len(s)}))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
