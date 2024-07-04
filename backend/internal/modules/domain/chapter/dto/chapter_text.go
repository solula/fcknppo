package dto

// ChapterText текст главы
type ChapterText struct {
	Uuid        string // Uuid
	ChapterUuid string // Uuid главы
	Text        string // Текст
}

type ChapterTexts []*ChapterText

// ChapterTextCreate модель создания текста главы
type ChapterTextCreate struct {
	ChapterUuid string // Uuid главы
	Text        string // Текст
}

// ChapterTextUpdate модель обновления текста главы
type ChapterTextUpdate struct {
	Text string // Текст
}
