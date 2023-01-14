package main

import (
	"log"

	bouyomichan "github.com/go-numb/go-bouyomichan"
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
	for i := 0; i < len(messages); i++ {
		client := bouyomichan.New("localhost:50001")
		if err := client.Speaking(messages[i]); err != nil {
			log.Fatal(err)
		}
	}
}
