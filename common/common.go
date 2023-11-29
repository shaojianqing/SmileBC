package common

const (
	HashLength    = 32
	AddressLength = 20
)

type Key []byte

type Code []byte

type Data []byte

type Content []byte

type Hash [HashLength]byte

type Address [AddressLength]byte

func HashEqual(src, dest Hash) bool {
	return false
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > AddressLength {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func (h *Hash) SetBytes(data []byte) {
	if len(data) > HashLength {
		data = data[len(data)-HashLength:]
	}

	copy(h[HashLength-len(data):], data)
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func LeftPadBytes(slice []byte, l int) []byte {
	if l < len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded[l-len(slice):], slice)

	return padded
}
