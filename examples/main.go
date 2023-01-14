package main

import (
	"time"

	"github.com/go-numb/go-bouyomichan"
)

var (
	messages = []string{
		"L知ってるか",
		"えげつねぇな",
		"おれがさん人分になる",
		"知らんけど",
		"ともだちってなんだろう",
		"キミたちは強くなる",
		"国なんてあてにしちゃダメ",
		"チャリでいく",
		"うそはウソであるとみぬける人でないとむつかしい",
		"だーりん、だいすきだっちゃ",
	}
)

func main() {
	bouyomi := bouyomichan.New()
	defer bouyomi.Close()

	for i := 0; i < len(messages); i++ {
		bouyomi.Speaking(messages[i])
		time.Sleep(time.Second)
	}

}
