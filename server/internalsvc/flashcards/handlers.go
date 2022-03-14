package flashcards

type DefaultHandler struct {
}

func (d DefaultHandler) DecodeCard(cards interface{}, cardType CardType) []FlashCard {
	panic("implement me")
}

func (d DefaultHandler) EncodeCard(Cards []FlashCard, cardType CardType) []byte {
	panic("implement me")
}

func (d DefaultHandler) ReadFlashcards(id uint64, cardType CardType) []FlashCard {
	panic("implement me")
}

func (d DefaultHandler) StoreCard(id uint64, Cards []FlashCard, cardType CardType) error {
	panic("implement me")
}
