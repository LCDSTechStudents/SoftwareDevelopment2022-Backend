package flashcards

type FlashCardHandler interface {
	DecodeCard(cards interface{}, cardType CardType) []FlashCard
	EncodeCard(Cards []FlashCard, cardType CardType) []byte

	ReadFlashcards(id uint64, cardType CardType) []FlashCard
	StoreCard(id uint64, Cards []FlashCard, cardType CardType) error
}

type FlashCard interface {
	GetContent() interface{}
}

func InitFlashCardHandler() FlashCardHandler {
	return &DefaultHandler{}
}
