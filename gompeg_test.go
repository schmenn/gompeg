package gompeg_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Schmenn/gompeg"
	"github.com/Schmenn/gompeg/types"
)

func TestMain(m *testing.M) {
	test, err := gompeg.Open("test.mp4")
	if err != nil {
		log.Fatalln(err)
	}

	test.SetOutputFile("test.webm")
	test.SetDimensions(200, types.AutoDimension)
	test.SetAudioSampleRate(types.SampleRate48k)
	test.SetStartTime(time.Second*3)
	test.SetEndTime(time.Second*8)

	fmt.Printf("%#v\n", test)
	
	err = test.Do()
	if err != nil {
		fmt.Println(err)
	}
}
