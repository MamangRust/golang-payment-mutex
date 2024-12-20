package repository

import (
	"fmt"
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/requests"
	recordmapper "payment-mutex/internal/mapper/record"
	"payment-mutex/internal/models"
	"payment-mutex/pkg/apikey"
	"strings"
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

func (ds *merchantRepository) ReadAll(page int, pageSize int, search string) ([]*record.MerchantRecord, int, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	filteredMerchants := make([]models.Merchant, 0)

	for _, merchant := range ds.merchants {
		if search == "" ||
			strings.Contains(strings.ToLower(merchant.Name), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(merchant.ApiKey), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(merchant.Status), strings.ToLower(search)) {
			filteredMerchants = append(filteredMerchants, merchant)
		}
	}

	totalRecords := len(filteredMerchants)

	start := (page - 1) * pageSize
	if start >= totalRecords {
		return nil, totalRecords, nil
	}

	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}

	paginatedMerchants := filteredMerchants[start:end]

	return ds.mapping.ToMerchantsRecord(paginatedMerchants), totalRecords, nil
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
