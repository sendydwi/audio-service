package audioservice_test

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sendydwi/audio-service/service/audioservice"
	mockaudioservice "github.com/sendydwi/audio-service/service/audioservice/mock"
	audioutils "github.com/sendydwi/audio-service/util/audio"
	mockstorageutils "github.com/sendydwi/audio-service/util/storage/mock"
)

const PATH_EXAMPlE_AUDIO = "./../../../example/input/sample-5.m4a"

func Test_UploadAudio_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockstorageutils.NewMockStorage(ctrl)
	mockRepo := mockaudioservice.NewMockRepository(ctrl)
	service := audioservice.AudioService{
		Repository: mockRepo,
		Storage:    mockStorage,
	}

	file, err := os.Open(PATH_EXAMPlE_AUDIO)
	if err != nil {
		log.Println("failed to load example audio")
	}

	gomock.InOrder(
		mockStorage.EXPECT().StoreFile("1/1", gomock.Any(), "wav", gomock.Any()).Return("filepath", nil).Times(1),
		mockRepo.EXPECT().Save("1", "1", "filepath").Times(1),
	)

	err = service.UploadAudioFile("1", "1", file)
	if err != nil {
		t.Fatal("upload should not return error", err)
	}
}

func Test_UploadAudio_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockstorageutils.NewMockStorage(ctrl)
	mockRepo := mockaudioservice.NewMockRepository(ctrl)
	service := audioservice.AudioService{
		Repository: mockRepo,
		Storage:    mockStorage,
	}

	file, err := os.Open("")
	if err != nil {
		log.Println("failed to load example audio")
	}

	err = service.UploadAudioFile("1", "1", file)
	if err == nil || err.Error() != "invalid argument" {
		t.Fatal("upload should return error invalid argument", err)
	}
}

func Test_GetAudio_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockstorageutils.NewMockStorage(ctrl)
	mockRepo := mockaudioservice.NewMockRepository(ctrl)
	service := audioservice.AudioService{
		Repository: mockRepo,
		Storage:    mockStorage,
	}

	file, err := os.Open(PATH_EXAMPlE_AUDIO)
	if err != nil {
		log.Println("failed to load example audio")
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("failed to read example audio")
	}

	gomock.InOrder(
		mockRepo.EXPECT().GetFilepath("1", "1").Return(PATH_EXAMPlE_AUDIO, nil).Times(1),
		mockStorage.EXPECT().GetFile(PATH_EXAMPlE_AUDIO).Return(*bytes.NewBuffer(data), nil).Times(1),
	)

	audioData, contentType, err := service.GetAudioFile("1", "1", "m4a")
	if err != nil {
		t.Fatal("get audio should not return error", err)
	}

	if len(audioData) == 0 {
		t.Fatal("get audio should not return empty data")
	}

	if contentType != audioutils.M4A.GetContentType() {
		t.Fatal("get audio suppose to return ", audioutils.M4A.GetContentType(), " but got ", contentType, " instead")
	}
}

func Test_GetAudio_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockstorageutils.NewMockStorage(ctrl)
	mockRepo := mockaudioservice.NewMockRepository(ctrl)
	service := audioservice.AudioService{
		Repository: mockRepo,
		Storage:    mockStorage,
	}

	errorMessage := "file with user id 1 and phrase id 1 not found"
	gomock.InOrder(
		mockRepo.EXPECT().GetFilepath("1", "1").Return("", errors.New(errorMessage)).Times(1),
	)

	_, _, err := service.GetAudioFile("1", "1", "m4a")
	if err == nil || err.Error() != errorMessage {
		t.Fatal("get audio should return error ", errorMessage)
	}

}
