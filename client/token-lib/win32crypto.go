package token_lib

import (
	"syscall"
	"unsafe"
)

const (
	CRYPTPTOTEXT_UI_FORBIDDEN = 0x1
)

var (
	dllcrypt32         = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32        = syscall.NewLazyDLL("Kernel32.dll")
	CryptUnprotectData = dllcrypt32.NewProc("CryptUnprotectData")
	LocalFree          = dllkernel32.NewProc("LocalFree")
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		cbData: uint32(len(d)),
		pbData: &d[0],
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB

	r, _, err := CryptUnprotectData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer LocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}

// https://stackoverflow.com/questions/33516053/windows-encrypted-rdp-passwords-in-golang
