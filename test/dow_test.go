package test

import (
	"github.com/Interesting-GO/youtubetools/video_dow"
	"log"
	"testing"
)

func TestDow(t *testing.T) {
	dow := video_dow.YoutubeDow("https://www.youtube.com/watch?v=NlVU4kpKLHI", "./test.mp4", "127.0.0.1:8001")

	if dow != nil {
		panic(dow)
	} else {
		log.Print("下载完毕")
	}
}
