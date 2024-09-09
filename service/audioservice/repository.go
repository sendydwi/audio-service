package audioservice

import (
	"github.com/sendydwi/audio-service/database/model"
	"gorm.io/gorm"
)

// create interface to make it easier to migrate
type Repository interface {
	Save(userId string, phraseId string, filepath string) error
	GetFilepath(userId string, phraseId string) (string, error)
}

type PosgresRepository struct {
	database *gorm.DB
}

func (p *PosgresRepository) Save(userId string, phraseId string, filepath string) error {
	audio := &model.Audio{
		UserId:    userId,
		PhraseId:  phraseId,
		AudioPath: filepath,
	}

	result := p.database.Create(audio)
	return result.Error
}

func (p *PosgresRepository) GetFilepath(userId string, phraseId string) (string, error) {
	var book model.Audio
	if result := p.database.First(&book, userId, phraseId); result.Error != nil {
		return "", result.Error
	}

	return book.AudioPath, nil
}
