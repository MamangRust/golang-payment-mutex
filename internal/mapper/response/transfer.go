package responseMapper

import (
	"payment-mutex/internal/domain/record"
	"payment-mutex/internal/domain/response"
)

type transferResponseMapper struct {
}

func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{}
}

func (s *transferResponseMapper) ToTransferResponse(transfer record.TransferRecord) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             transfer.TransferID,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
	}
}

func (s *transferResponseMapper) ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponse(*response))
	}

	return responses
}
