package bouyomichan

import (
	"log"
	"testing"
)

var msgs = []string{
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

func TestSpeeking(t *testing.T) {
	client := New("localhost:50001")

	for i, v := range msgs {
		switch i {
		case 1:
			client.Voice = VoiceMan01

		case 2:
			client.Voice = VoiceNeutral
		}

		if err := client.Speaking(v); err != nil {
			log.Fatal(err)
		}
	}
}
