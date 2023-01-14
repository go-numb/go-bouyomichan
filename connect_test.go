package bouyomichan

import (
	"log"
	"testing"
)

var msgs = []string{
	"すたーと",
	"ボリューム下げておとこのひとに",
	"取りつかれている",
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

	// client.Voice = VoiceRobot01
	// if err := client.Speaking("こえをろぼっとに"); err != nil {
	// 	t.Fatal(err)
	// }

	// client.Voice = VoiceNeutral
	// if err := client.Speaking("こえをナチュラルに"); err != nil {
	// 	t.Fatal(err)
	// }

	// client.Voice = VoiceWoman01
	// if err := client.Speaking("戻してターンエンド"); err != nil {
	// 	t.Fatal(err)
	// }

	// var (
	// 	i     = 0
	// 	count = 10
	// )
	// ticker := time.NewTicker(3 * time.Second)

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		if err := client.Speaking(fmt.Sprintf("%d", i)); err != nil {
	// 			t.Fatal(err)
	// 		}
	// 		i++
	// 		if i > count {
	// 			return
	// 		}
	// 	}
	// }
}
