package audioservice

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/iFaceless/godub/converter"
	audioutils "github.com/sendydwi/audio-service/util/audio"
	storageutils "github.com/sendydwi/audio-service/util/storage"
)

type AudioService struct {
	Repository Repository
	Storage    storageutils.Storage
}

func (a AudioService) UploadAudioFile(userId, phraseId string, audioFile multipart.File) error {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	err := converter.NewConverter(writer).WithDstFormat(audioutils.WAV.String()).Convert(audioFile)
	if err != nil {
		log.Println("failed to convert audio file")
		return err
	}

	targetFolder := fmt.Sprintf("%s/%s", userId, phraseId)
	path, err := a.Storage.StoreFile(targetFolder, uuid.NewString(), audioutils.WAV.String(), buf)
	if err != nil {
		log.Println("failed to store audio file")
		return err
	}

	err = a.Repository.Save(userId, phraseId, path)
	if err != nil {
		log.Println("failed to write data to database")
		return err
	}
	return nil
}

func (a AudioService) GetAudioFile(userId, phraseId, audioFormat string) ([]byte, string, error) {
	path, err := a.Repository.GetFilepath(userId, phraseId)
	if err != nil {
		return nil, "", err
	}

	targetFormat := audioutils.GetSupportedAudioFormatByString(audioFormat)
	if targetFormat == audioutils.NotSupported {
		return nil, "", errors.New("target audio format not supported")
	}

	filedata, err := a.Storage.GetFile(path)
	if err != nil {
		return nil, "", err
	}

	if targetFormat == audioutils.WAV {
		return filedata.Bytes(), targetFormat.GetContentType(), nil
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	err = converter.NewConverter(writer).WithDstFormat(targetFormat.String()).Convert(bytes.NewBuffer(filedata.Bytes()))
	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), targetFormat.GetContentType(), nil
}
