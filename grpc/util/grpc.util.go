package grpcutil

import (
	"math"

	"auth/proto/gcommon"
)

func AsPageMetadata(pageable *gcommon.Pageable, totalElements int) *gcommon.PageMetadata {
	if pageable.PagingIgnored {
		return &gcommon.PageMetadata{
			Page:          0,
			Size:          int32(totalElements),
			TotalElements: int64(totalElements),
			TotalPages:    1,
			HasNext:       false,
			HasPrevious:   false,
		}
	}

	totalPages := int32(math.Ceil(float64(totalElements) / float64(pageable.Size)))
	hasPrevious := pageable.Page > 0
	hasNext := totalPages > (pageable.Page + 1)
	return &gcommon.PageMetadata{
		Page:          pageable.Page,
		Size:          pageable.Size,
		TotalElements: int64(totalElements),
		TotalPages:    totalPages,
		HasNext:       hasNext,
		HasPrevious:   hasPrevious,
	}
}

func AsEmptyResponse() *gcommon.EmptyResponse {
	return &gcommon.EmptyResponse{}
}
