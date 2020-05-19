package service

import (
	"discord-slot/model"
	"discord-slot/repository"
	"math/rand"
)

type EmojiService struct {
	er repository.EmojiRepository
}

func (s *EmojiService) RegisterEmoji(emoji model.Emoji) {
	s.er.Save(emoji)
}

func (s *EmojiService) GetAllEmoji() []model.Emoji {
	return s.er.GetAll()
}

func (s *EmojiService) GetRandomEmojis(n int) []model.Emoji {
	src := s.er.GetAll()
	retVal := make([]model.Emoji, n)
	for i := 0; i < n; i++ {
		emoji := src[rand.Intn(len(src))]
		retVal = append(retVal, emoji)
	}
	return retVal
}

func (s *EmojiService) DeleteEmoji(emoji model.Emoji) {
	s.er.Delete(emoji)
}
