package model

type Audio struct {
	UserId    string `json:"user_id" gorm:"primaryKey"`
	PhraseId  string `json:"phrase_id" gorm:"primaryKey"`
	AudioPath string `json:"audio_path"`
}
