package gompeg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/Schmenn/gompeg/types"
)




// Command gets executed with Do()
type Command struct {
	inputFile       string
	OutputFile      string
	Framerate       uint8
	VideoCodec      types.VideoCodec
	AudioCodec      string
	audioSampleRate types.AudioSampleRate
	Dimensions      types.Dimensions
	audioDisabled   bool
	startTime       time.Duration
	endTime         time.Duration
}


// Open creates a Command struct
func Open(file string) (*Command, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, errors.New("File not found")
	}
	return &Command{
		inputFile: file,
	}, nil
}

// SetFrameRate sets the framerate
func (c *Command) SetFrameRate(r int) *Command {
	c.Framerate = uint8(r)
	return c
}

// SetOutputFile sets the output file
func (c *Command) SetOutputFile(file string) *Command {
	c.OutputFile = file
	return c
}

// SetVideoCodec sets the output video codec
func (c *Command) SetVideoCodec(vc types.VideoCodec) *Command {
	c.VideoCodec = vc
	return c
}

// SetDimensions sets the output video dimensions
func (c *Command) SetDimensions(x int, y int) *Command {
	c.Dimensions.X = x
	c.Dimensions.Y = y
	return c
}

// DisableAudio disables audio
func (c *Command) DisableAudio() *Command {
	c.audioDisabled = true
	return c
}

// SetAudioSampleRate sets audio sample rate
func (c *Command) SetAudioSampleRate(ab types.AudioSampleRate) *Command {
	c.audioSampleRate = ab
	return c
}

// SetStartTime sets the time at which the video starts
func (c *Command) SetStartTime(start time.Duration) *Command {
	c.startTime = start
	return c
}

// SetEndTime sets the time at which the video ends
func (c *Command) SetEndTime(end time.Duration) *Command {
	c.endTime = end
	return c
}

// Do executes the command
func (c *Command) Do() error {
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		return errors.New("FFmpeg was not found in the path")
	}
	var args []string
	args = append(args, "-i", c.inputFile)
	if c.Framerate != 0 {
		args = append(args, "-r", strconv.Itoa(int(c.Framerate)))
	}
	if c.VideoCodec != "" {
		args = append(args, "-c:v", string(c.VideoCodec))
	}
	if c.Dimensions != (types.Dimensions{}) {
		args = append(args, "-vf", "scale="+strconv.Itoa(c.Dimensions.X)+":"+strconv.Itoa(c.Dimensions.Y))
	}
	if c.audioDisabled {
		args = append(args, "-an")
	}
	if c.audioSampleRate != 0 {
		args = append(args, "-b:a", strconv.Itoa(int(c.audioSampleRate)))
	}
	if c.startTime != 0 {
		args = append(args, "-ss", ffmpegTimeStampFromDuration(c.startTime))
	}
	if c.endTime != 0 {
		args = append(args, "-to", ffmpegTimeStampFromDuration(c.endTime))
	}

	// at the end
	if c.OutputFile == "" {
		return errors.New("No Output file provided")
	}
	args = append(args, c.OutputFile, "-y")

	cmd := exec.Command(ffmpeg, args...)
	fmt.Println(cmd.String())
	/*out, err := cmd.CombinedOutput()
	fmt.Println(string(out))*/
	return cmd.Run()
}

func ffmpegTimeStampFromDuration(in time.Duration) string {
	H := in / time.Hour
	in -= H * time.Hour
	M := in / time.Minute
	in -= M * time.Minute
	S := in / time.Second
	in -= S * time.Second
	MS := in / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d.%03d", H, M, S, MS)
}
