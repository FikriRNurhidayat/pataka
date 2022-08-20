package feature

import (
	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"github.com/fikrirnurhidayat/ffgo/internal/manager"
)

type CreateParams struct {
	Name    string
	Label   string
	Enabled bool
}

type CreateResult struct {
	Feature *domain.Feature
}

type ListParams struct {
	manager.PaginationParams
	Sort    string
	Q       string
	Enabled bool
}

type ListResult struct {
	manager.PaginationResult
	Size     uint32
	Features []domain.Feature
}

type GetParams struct {
	Name string
}

type GetResult struct {
	Feature *domain.Feature
}

type UpdateParams struct {
	Name    string
	Label   string
	Enabled bool
}

type UpdateResult struct {
	Feature *domain.Feature
}

type DeleteParams struct {
	Name string
}
