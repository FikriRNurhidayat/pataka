package transaction

import (
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
)

type RepositoryFactory struct {
	FeatureRepository  driver.Factory[domain.FeatureRepository]
	AudienceRepository driver.Factory[domain.AudienceRepository]
}
