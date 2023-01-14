# 棒読みちゃんに詠んでもらおうネット名言

## Require
- Go
- [棒読みちゃん](https://chi.usamimi.info/Program/Application/BouyomiChan/)

## Usage
```go
package main

import (
    "time"
	"bouyomichan"
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
	client := bouyomichan.New("localhost:50001")

	for i, v := range messages {
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
```


おす！

