package service

import (
	"discord-slot/model"
	"discord-slot/repository"
)

type EmojiRegisterService struct {
	er repository.EmojiRepository
}

func (s *EmojiRegisterService) RegisterEmoji(emoji model.Emoji) {
	s.er.Save(emoji)
}
