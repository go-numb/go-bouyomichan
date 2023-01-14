package bouyomichan

import (
	"encoding/binary"
	"log"
	"net"
)

type Client struct {
	conn net.Conn
}

func New() *Client {
	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		conn: conn,
	}
}

func (p *Client) Close() error {
	return p.conn.Close()
}

// Speaking ローカルなりに棒読みちゃんアプリが立ち上げられ、かつ、設定のアプリケーション連携がTrue担っていることが前提
func (p *Client) Speaking(msg string) error {
	if _, err := p.conn.Write(
		verbalizing(
			[]byte(msg),
		),
	); err != nil {
		return err
	}

	return nil
}

// verbalizing tcp websocket用byteを作る
// 各パラメータ引数未設定
func verbalizing(m []byte) []byte {
	msg := string(m)
	msg_length := uint32(len(msg))
	iCommand := []byte{1, 0}
	iSpeed := []byte{100, 0}
	iTone := []byte{255, 255}
	iVolume := []byte{255, 255}
	iVoice := []byte{0, 0}
	bCode := []byte{0}

	d := append(iCommand, iSpeed...)
	d = append(d, iTone...)
	d = append(d, iVolume...)
	d = append(d, iVoice...)
	d = append(d, bCode...)

	bMsglength := make([]byte, 4)
	binary.LittleEndian.PutUint32(bMsglength, msg_length)

	d = append(d, bMsglength...)
	d = append(d, msg...)

	return d
}
