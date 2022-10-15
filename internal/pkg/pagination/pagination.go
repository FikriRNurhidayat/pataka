package pagination

import (
	"math"
)

type PaginationParams struct {
	PageNumber uint32
	PageSize   uint32
}

type PaginationResult struct {
	PageNumber uint32
	PageSize   uint32
	PageCount  uint32
	Size       uint32
}

func (params *PaginationParams) GetLimit() uint32 {
	if params.PageSize == 0 {
		return 10
	}

	return params.PageSize
}

func (params *PaginationParams) GetOffset() uint32 {
	if params.PageNumber == 0 {
		params.PageNumber = 1
	}

	return (params.PageNumber - 1) * params.PageSize
}

func (params *PaginationParams) PaginationResult(size uint32) *PaginationResult {
	result := &PaginationResult{}

	result.PageNumber = params.PageNumber
	result.PageSize = params.PageSize
	result.PageCount = uint32(math.Ceil(float64(size) / float64(result.PageSize)))
	result.Size = size

	return result
}
