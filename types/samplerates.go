package types

// AudioSampleRate Audio Sample Rate
type AudioSampleRate uint32

// consts
const (
	AutoDimension int = -1

	// Audio Sample Rates
	SampleRate8k     AudioSampleRate = 8000
	SampleRate11025  AudioSampleRate = 11025
	SampleRate16k    AudioSampleRate = 16000
	SampleRate22050  AudioSampleRate = 22050
	SampleRate44100  AudioSampleRate = 44100
	SampleRate48k    AudioSampleRate = 48000
	SampleRate88200  AudioSampleRate = 88200
	SampleRate96k    AudioSampleRate = 96000
	SampleRate176400 AudioSampleRate = 176400
	SampleRate192k   AudioSampleRate = 192000
	SampleRate352800 AudioSampleRate = 352800
	SampleRate384000 AudioSampleRate = 384000
)
