package bouyomichan

import (
	"log"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
)

var msgs = []string{
	`ミンナ💑で、おじサン🙋‍♂️構文✏️💕の、練習⛳️を、シタ💪😅💦んだケド‼️⁉️‼️ななチャン🙋‍♀️、優しくて😍💕❤️、惚れられちゃっタ😻💓カナ⁉️⁉️ホテル🏩💖トカ、誘われちゃタら😻❣️💕⁉️、どうしよう🥺‼️💓⁉️💖ナンチャッテ😅💦（笑）#おじさん構文`,
	"えげつねぇな",
	"おれがおじさん構文になる",
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

func TestScraping(t *testing.T) {
	client := New("localhost:50001")
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div.t_b", func(e *colly.HTMLElement) {
		if err := client.Speaking(e.Text); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://itainews.com/archives/2021300.html")
}
