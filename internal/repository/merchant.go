package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"payment-mutex/pkg/apikey"
	"sync"
)

type merchantRepository struct {
	mu        sync.RWMutex
	merchants map[int]models.Merchant
	nextID    int
	mapping   recordmapper.MerchantRecordMapping
}

func NewMerchantRepository(mapping recordmapper.MerchantRecordMapping) *merchantRepository {
	return &merchantRepository{
		merchants: make(map[int]models.Merchant),
		nextID:    1,
		mapping:   mapping,
	}
}

func (ds *merchantRepository) ReadAll() ([]*record.MerchantRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	merchants := make([]models.Merchant, 0, len(ds.merchants))

	for _, merchant := range ds.merchants {
		merchants = append(merchants, merchant)
	}

	if len(merchants) == 0 {
		return nil, fmt.Errorf("no merchant found")
	}

	return ds.mapping.ToMerchantsRecord(merchants), nil

}

func (ds *merchantRepository) Read(merchantID int) (*record.MerchantRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	merchant, ok := ds.merchants[merchantID]

	if !ok {
		return nil, fmt.Errorf("merchant with ID %d not found", merchantID)
	}

	return ds.mapping.ToMerchantRecord(merchant), nil
}

func (ds *merchantRepository) ReadByName(name string) (*record.MerchantRecord, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, merchant := range ds.merchants {
		if merchant.Name == name {
			return ds.mapping.ToMerchantRecord(merchant), nil
		}
	}

	return nil, fmt.Errorf("merchant not found")
}

func (ds *merchantRepository) ReadByApiKey(apiKey string) (*record.MerchantRecord, error) {
	ds.mu.RLock()

	defer ds.mu.RUnlock()

	for _, merchant := range ds.merchants {
		if merchant.ApiKey == apiKey {
			return ds.mapping.ToMerchantRecord(merchant), nil
		}
	}

	return nil, fmt.Errorf("merchant not found")
}

func (ds *merchantRepository) Create(request requests.CreateMerchantRequest) (*record.MerchantRecord, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	merchant := models.Merchant{
		MerchantID: ds.nextID,
		ApiKey:     apikey.GenerateApiKey(),
		Name:       request.Name,
		UserID:     request.UserID,
		Status:     "active",
	}

	ds.merchants[merchant.MerchantID] = merchant
	ds.nextID++

	return ds.mapping.ToMerchantRecord(merchant), nil
}

func (ds *merchantRepository) Update(request requests.UpdateMerchantRequest) (*record.MerchantRecord, error) {
	ds.mu.Lock()

	defer ds.mu.Unlock()

	merchant, ok := ds.merchants[request.MerchantID]

	if !ok {
		return nil, fmt.Errorf("merchant with id %d not found", request.MerchantID)
	}

	merchant.Name = request.Name
	merchant.UserID = request.UserID
	merchant.Status = request.Status

	ds.merchants[request.MerchantID] = merchant

	return ds.mapping.ToMerchantRecord(merchant), nil

}

func (ds *merchantRepository) Delete(merchantID int) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.merchants[merchantID]; ok {
		delete(ds.merchants, merchantID)

		return nil
	}

	return fmt.Errorf("merchant with id %d not found", merchantID)
}
