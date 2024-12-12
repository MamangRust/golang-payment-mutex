package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type merchantResponseMapper struct {
}

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

func (s *merchantResponseMapper) ToMerchantResponse(merchant record.MerchantRecord) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:     merchant.MerchantID,
		Name:   merchant.Name,
		Status: merchant.Status,
		ApiKey: merchant.ApiKey,
	}
}

func (s *merchantResponseMapper) ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse {
	var response []*response.MerchantResponse
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponse(*merchant))
	}
	return response
}
