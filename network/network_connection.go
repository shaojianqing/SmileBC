package p2p

import (
	"encoding/json"
	"errors"
	"net"
)

const (
	MaxMessageSize = ^uint32(0) >> 8
)

type MessageReader interface {
	Read() (*Message, error)
}

type MessageWriter interface {
	Write(*Message) error
}

type MessageConn interface {
	MessageReader
	MessageWriter
}

type MConnection struct {
	conn   net.Conn
	secret *Secret
}

func NewMConnection(conn net.Conn, secret *Secret) MessageConn {
	return &MConnection{
		conn:   conn,
		secret: secret,
	}
}

func (conn *MConnection) Read() (*Message, error) {
	buffer := make([]byte, MaxMessageSize)
	count, err := conn.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	cipher := buffer[:count]
	plain := conn.secret.Decrypt(cipher)

	message := &Message{}
	err = json.Unmarshal(plain, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (conn *MConnection) Write(message *Message) error {
	plain, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if len(plain) > int(MaxMessageSize) {
		return errors.New("message size is too large than the maximum limit")
	}

	cipher := conn.secret.Encrypt(plain)
	_, err = conn.conn.Write(cipher)
	if err != nil {
		return err
	}
	return nil
}
