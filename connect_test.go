package bouyomichan

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
)

var msgs = []string{
	"えげつねぇな",
	`1.適切な衣服の着用
	首筋が冷える原因として、適切な衣服を着用していないことが挙げられます。冬場は防寒着を着用し、暖かいマフラーや襟巻きなどを使用することで首筋の冷えを防ぐことができます。
	
	2.適度な運動
	運動不足が首筋の冷えを引き起こすことがあります。適度な運動を行うことで、全身の血行が良くなり、首筋の冷えを予防することができます。ただし、冷え性の人は運動前後に十分なストレッチを行うことが重要です。
	
	3.食事の改善
	冷え性には、体を温める効果のある食材を積極的に摂取することも重要です。生姜やにんにく、唐辛子などのスパイスやハーブを使った料理、また、温かいスープや鍋料理などが、首筋の冷え対策に役立ちます。
	
	4.ストレスの軽減
	ストレスが長期的に続くと自律神経のバランスが崩れ、血行不良が引き起こされることがあります。ストレスを軽減するためには、定期的なリラックスタイムの確保や趣味やスポーツなどを通じた適度な運動などが効果的です。
	
	以上のような対処法を実践することで、首筋の冷えを予防することができます。しかし、首筋の冷えが長期的に続く場合には、病気や疾患の可能性もあるため、定期的な医療機関の受診が必要です`,
	"恐ろしく早いしゅとう、オレじゃなきゃ見逃しちゃうね",
}

func TestSpeeking(t *testing.T) {
	client := New("localhost:50001")

	client.Volume = 100

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

		// Timeoutのテスト
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

	L:
		for {
			select {
			case <-ctx.Done():
				if err := client.Stop(); err != nil {
					fmt.Println(err)
					break L
				}
				break L
			default:
				if client.IsNowPlayng() {
					fmt.Println("wait...")
					time.Sleep(time.Second)
					continue
				}
				break L
			}
		}
		fmt.Println("next")
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
