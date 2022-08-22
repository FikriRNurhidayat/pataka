package pagination

import "math"

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

func (params *PaginationParams) ToLimit(defaultPageSize uint32) uint32 {
	if defaultPageSize == 0 {
		defaultPageSize = 10
	}

	if params.PageSize == 0 {
		return defaultPageSize
	}

	return params.PageSize
}

func (params *PaginationParams) ToOffset(defaultPageNumber uint32) uint32 {
	if defaultPageNumber == 0 {
		defaultPageNumber = 1
	}

	if params.PageNumber == 0 {
		params.PageNumber = defaultPageNumber
	}

	return (params.PageNumber - 1) * params.PageSize
}

func (params *PaginationParams) ToPaginationResult(size uint32) *PaginationResult {
	result := &PaginationResult{}

	if params.PageNumber == 0 {
		result.PageNumber = 1
	}

	if params.PageSize == 0 {
		result.PageSize = 10
	}

	result.PageCount = uint32(math.Ceil(float64(size) / float64(result.PageSize)))
	result.Size = size

	return result
}
