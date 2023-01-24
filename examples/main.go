package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-numb/go-bouyomichan"
	"github.com/gocarina/gocsv"
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
	// UseCSV()
	// return
	client := bouyomichan.New("localhost:50001")
	for i := 0; i < len(messages); i++ {
		if i%2 == 0 {
			client.Voice = bouyomichan.VoiceDefault
			client.Tone = 105
		} else {
			client.Voice = bouyomichan.VoiceNeutral
			client.Tone = 110
		}

		if err := client.Speaking(messages[i]); err != nil {
			log.Fatal(err)
		}
	}
}

func UseCSV() {
	q := ""
	f, _ := os.Open(fmt.Sprintf(`./posts-%s.csv`, q))
	defer f.Close()

	r := gocsv.DefaultCSVReader(f)

	str, _ := r.ReadAll()
	fmt.Printf("%#v", str)

	for i, values := range str {
		for j, v := range values {
			if j == 5 {
				client := bouyomichan.New("localhost:50001")
				if err := client.Speaking(v); err != nil {
					log.Fatal(err)
				}
			}
		}

		if i > 5 {
			return
		}
	}
}
