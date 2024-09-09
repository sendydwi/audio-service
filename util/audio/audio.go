package audioutils

import "strings"

type SupportedAudioFormat int

const (
	NotSupported SupportedAudioFormat = iota
	M4A
	MP3
	WAV
)

func (f SupportedAudioFormat) String() string {
	list := []string{"notsupported", "m4a", "mp3", "wav"}
	return list[f]
}

func (f SupportedAudioFormat) GetContentType() string {
	list := []string{"notsupported", "audio/mp4", "audio/mpeg", "audio/wav"}
	return list[f]
}

func GetSupportedAudioFormatByString(format string) SupportedAudioFormat {
	format = strings.ToLower(format)
	switch format {
	case "mp3":
		return MP3
	case "m4a":
		return M4A
	case "wav":
		return WAV
	default:
		return NotSupported
	}
}
