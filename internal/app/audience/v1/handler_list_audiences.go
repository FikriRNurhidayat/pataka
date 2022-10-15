package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) ListAudiences(ctx context.Context, req *audiencev1.ListAudiencesRequest) (*audiencev1.ListAudiencesResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.ReadAudienceScope); err != nil {
		return nil, err
	}

	var (
		enabled *bool
		e       bool
	)

	switch req.GetStatus() {
	case audiencev1.AudienceStatus_AUDIENCE_STATUS_UNSPECIFIED:
		enabled = nil
	case audiencev1.AudienceStatus_AUDIENCE_STATUS_ENABLED:
		e = true
		enabled = &e
	case audiencev1.AudienceStatus_AUDIENCE_STATUS_DISABLED:
		e = false
		enabled = &e
	}

	result, err := s.listAudiencesService.Call(ctx, &domain.ListAudiencesParams{
		PaginationParams: &pagination.PaginationParams{
			PageNumber: req.GetPageNumber(),
			PageSize:   req.GetPageSize(),
		},
		Sort:        req.GetSort(),
		FeatureName: req.GetFeatureName(),
		AudienceIds: req.GetAudienceId(),
		Enabled:     enabled,
	})
	if err != nil {
		return nil, err
	}

	return ToAudiencesProtoResponse(result), nil
}
