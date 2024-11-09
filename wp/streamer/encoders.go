package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

type VideoEncoder struct {
}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// Create a transcoder
	trans := new(transcoder.Transcoder)

	// Build the output path
	output := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)

	// Initialize the transcoder
	err := trans.Initialize(v.InputFile, output)
	if err != nil {
		return err
	}
	// Set codec
	trans.MediaFile().SetVideoCodec("libx264")

	// Start transcoding processs
	done := trans.Run(false)

	err = <-done
	if err != nil {
		return err
	}

	return nil
}
