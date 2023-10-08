package common

const (
	HashLength    = 32
	AddressLength = 20
)

type Code []byte

type Content []byte

type Hash [HashLength]byte

type Address [AddressLength]byte
