package audience

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/pagination"
	audiencev1 "github.com/fikrirnurhidayat/ffgo/protobuf/audience/v1"
)

func (s *Server) CreateAudience(ctx context.Context, req *audiencev1.CreateAudienceRequest) (*audiencev1.CreateAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteAudienceScope); err != nil {
		return nil, err
	}

	result, err := s.createAudienceService.Call(ctx, &domain.CreateAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.CreateAudienceResponse](result.Audience), nil
}

func (s *Server) GetAudience(ctx context.Context, req *audiencev1.GetAudienceRequest) (*audiencev1.GetAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.ReadAudienceScope); err != nil {
		return nil, err
	}

	result, err := s.getAudienceService.Call(ctx, &domain.GetAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.GetAudienceResponse](result.Audience), nil
}

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

func (s *Server) UpdateAudience(ctx context.Context, req *audiencev1.UpdateAudienceRequest) (*audiencev1.UpdateAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteAudienceScope); err != nil {
		return nil, err
	}

	result, err := s.updateAudienceService.Call(ctx, &domain.UpdateAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
		Enabled:     req.GetEnabled(),
	})
	if err != nil {
		return nil, err
	}

	return ToAudienceProtoResponse[audiencev1.UpdateAudienceResponse](result.Audience), nil
}

func (s *Server) DeleteAudience(ctx context.Context, req *audiencev1.DeleteAudienceRequest) (*audiencev1.DeleteAudienceResponse, error) {
	if _, err := s.authenticationService.Call(ctx, domain.WriteAudienceScope); err != nil {
		return nil, err
	}

	err := s.deleteAudienceService.Call(ctx, &domain.DeleteAudienceParams{
		AudienceId:  req.GetAudienceId(),
		FeatureName: req.GetFeatureName(),
	})
	if err != nil {
		return nil, err
	}

	return &audiencev1.DeleteAudienceResponse{}, nil
}
