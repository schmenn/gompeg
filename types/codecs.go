package types

// VideoCodec string alias
type VideoCodec string

// Video and Audio codecs
const (
	Copy VideoCodec = "copy"

	VP8 VideoCodec = "libvpx"
	VP9 VideoCodec = "libvpx-vp9"
	AMV  VideoCodec = "amv"
	APNG VideoCodec = "apng"
	AV1  VideoCodec = "av1"
	Bmp  VideoCodec = "bmp"
	FLV  VideoCodec = "flv1"
	GIF  VideoCodec = "gif"
	H261 VideoCodec = "h261"
	H263 VideoCodec = "h263"
	H264 VideoCodec = "h264"
	H265 VideoCodec = "hevc"
	MP4  VideoCodec = "mpeg4"
)
