package repository

import "discord-slot/model"

type EmojiRepository interface {
	Save(emoji model.Emoji)
	GetAll() []model.Emoji
	Delete(emoji model.Emoji)
}

type InMemoryEmojiRepository struct {
	emojis []model.Emoji
}

func (r *InMemoryEmojiRepository) Save(emoji model.Emoji) {
	for _, v := range r.emojis {
		if v == emoji {
			return
		}
	}
	r.emojis = append(r.emojis, emoji)
}

func (r *InMemoryEmojiRepository) GetAll() []model.Emoji {
	return r.emojis
}

func (r *InMemoryEmojiRepository) Delete(emoji model.Emoji) {
	emojis := r.emojis
	for i, v := range emojis {
		if v == emoji {
			emojis = append(emojis[:i], emojis[i+1:]...)
			new := make([]model.Emoji, len(emojis))
			r.emojis = new
		}
	}
}

func NewInMemoryEmojiRepository() InMemoryEmojiRepository {
	return InMemoryEmojiRepository{[]model.Emoji{}}
}
