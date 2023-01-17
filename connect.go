package bouyomichan

import (
	"encoding/binary"
	"log"
	"net"
)

type Command int16

const (
	CommandStart Command = 1
	CommandPause Command = iota * 1 * 16
	CommandResume
	CommandSkip
	CommandClear
)

type Voices int16

const (
	// Voice 1:女性1、2:女性2、3:男性1、4:男性2、5:中性、6:ロボット、7:機械1、8:機械2、10001～:SAPI5）
	VoiceDefault Voices = iota
	VoiceWoman01
	VoiceWoman02
	VoiceMan01
	VoiceMan02
	VoiceNeutral
	VoiceRobot01
	VoiceRobot02
	VoiceRobot03
	VoiceLocal
)

type Code int8

// Code... 文字コード
const (
	CodeUTF8 Code = iota
	CodeUnicode
	CodeShiftJIS
)

type Client struct {
	addr string

	Speed  int16
	Tone   int16
	Volume int16
	Voice  Voices
	Code   Code
}

func New(addr string) *Client {
	return &Client{
		addr: addr,
		// Default
		Speed:  100,
		Tone:   -1,
		Volume: 20,
		Voice:  VoiceDefault,
		Code:   CodeUTF8,
	}
}

// Speaking ローカルなりに棒読みちゃんアプリが立ち上げられ、かつ、設定のアプリケーション連携がTrue担っていることが前提
// 棒読みちゃんからEOFが返ってくるため常々接続
func (p *Client) Speaking(msg string) error {
	conn, err := net.Dial("tcp", p.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(
		verbalizing(
			[]byte(msg), p.Speed, p.Tone, p.Volume, p.Voice, p.Code,
		),
	)
	if err != nil {
		log.Println("[ERROR] ", err)
		return err
	}

	// READ EOF
	var buf = make([]byte, 16)
	n, _ := conn.Read(buf)
	if n != 0 {
		log.Println("[INFO] READ: ", n, buf)
	}

	return nil
}

// verbalizing tcp websocket用byteを作る
// 各パラメータ引数未設定
func verbalizing(message []byte, speed, tone, volume int16, voice Voices, bcode Code) []byte {
	msg := string(message)
	msg_length := uint32(len(msg))

	// var iCommand int16 = int16(1)
	// var d = make([]byte, 2)
	// binary.BigEndian.PutUint16(d, uint16(iCommand))
	// log.Println("command: ", d)

	d := []byte{1, 0}

	var tmp = make([]byte, 2)
	if speed != -1 && speed < 50 || speed > 300 {
		speed = -1
	}
	binary.LittleEndian.PutUint16(tmp, uint16(speed))
	d = append(d, tmp...)

	if tone != -1 && tone < 50 || tone > 200 {
		tone = -1
	}
	binary.LittleEndian.PutUint16(tmp, uint16(tone))
	d = append(d, tmp...)

	if volume != -1 && volume < -1 || volume > 100 {
		volume = -1
	}
	binary.LittleEndian.PutUint16(tmp, uint16(volume))
	d = append(d, tmp...)

	binary.LittleEndian.PutUint16(tmp, uint16(voice))
	d = append(d, tmp...)

	binary.BigEndian.PutUint16(tmp, uint16(bcode))
	d = append(d, tmp[1])

	bMsglength := make([]byte, 4)
	binary.LittleEndian.PutUint32(bMsglength, msg_length)

	d = append(d, bMsglength...)
	d = append(d, []byte(msg)...)

	return d
}
