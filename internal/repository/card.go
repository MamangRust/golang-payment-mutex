package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"payment-mutex/pkg/randomvcc"
	"strconv"
	"sync"
)

type cardRepository struct {
	mu      sync.RWMutex
	cards   map[int]models.Card
	nextID  int
	mapping recordmapper.CardRecordMapping
}

func NewCardRepository(mapping recordmapper.CardRecordMapping) *cardRepository {
	return &cardRepository{
		cards:   make(map[int]models.Card),
		nextID:  1,
		mapping: mapping,
	}
}

func (ds *cardRepository) ReadAll() ([]*record.CardRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	cards := make([]models.Card, 0, len(ds.cards))

	for _, card := range ds.cards {
		cards = append(cards, card)
	}

	if len(cards) == 0 {
		return nil, fmt.Errorf("no card found")
	}

	return ds.mapping.ToCardsRecord(cards), nil
}

func (ds *cardRepository) Read(cardID int) (*record.CardRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	card, ok := ds.cards[cardID]

	if !ok {
		return nil, fmt.Errorf("card with ID %d not found", cardID)
	}

	return ds.mapping.ToCardRecord(card), nil
}

func (ds *cardRepository) ReadByCardNumber(cardNumber string) (*record.CardRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	for _, card := range ds.cards {
		if card.CardNumber == cardNumber {
			return ds.mapping.ToCardRecord(card), nil
		}
	}

	return nil, fmt.Errorf("card with card number %s not found", cardNumber)
}

func (ds *cardRepository) ReadByUserID(userID int) (*record.CardRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	for _, card := range ds.cards {
		if card.UserID == userID {
			return ds.mapping.ToCardRecord(card), nil
		}
	}

	return nil, fmt.Errorf("card for user ID %d not found", userID)

}

func (ds *cardRepository) ReadByUsersID(userID int) ([]*record.CardRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	var userCards []models.Card

	for _, card := range ds.cards {
		if card.UserID == userID {
			userCards = append(userCards, card)
		}
	}

	if len(userCards) == 0 {
		return nil, fmt.Errorf("no card found for user ID %d", userID)
	}

	return ds.mapping.ToCardsRecord(userCards), nil
}

func (ds *cardRepository) Create(request requests.CreateCardRequest) (*record.CardRecord, error) {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	for _, existingCard := range ds.cards {
		if existingCard.UserID == request.UserID {
			return nil, fmt.Errorf("card for user ID %d already exists", request.UserID)
		}
	}

	random, err := randomvcc.RandomCardNumber()

	if err != nil {
		return nil, fmt.Errorf("random vcc error: %d", err)
	}

	card := models.Card{
		CardID:       ds.nextID,
		UserID:       request.UserID,
		CardNumber:   strconv.Itoa(int(random)),
		CardType:     request.CardType,
		ExpireDate:   request.ExpireDate,
		CVV:          request.CVV,
		CardProvider: request.CardProvider,
	}

	ds.cards[card.CardID] = card
	ds.nextID++

	return ds.mapping.ToCardRecord(card), nil
}

func (ds *cardRepository) Update(request requests.UpdateCardRequest) (*record.CardRecord, error) {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	card, ok := ds.cards[request.CardID]

	if !ok {
		return nil, fmt.Errorf("card with ID %d not found", request.CardID)
	}

	card.CardType = request.CardType
	card.ExpireDate = request.ExpireDate
	card.CVV = request.CVV
	card.CardProvider = request.CardProvider

	ds.cards[card.CardID] = card

	return ds.mapping.ToCardRecord(card), nil
}

func (ds *cardRepository) Delete(cardID int) error {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	if _, ok := ds.cards[cardID]; ok {
		delete(ds.cards, cardID)
		return nil
	}

	return fmt.Errorf("card with ID %d not found", cardID)
}
