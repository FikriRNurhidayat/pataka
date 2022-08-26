package domain

type Repository interface {
	FeatureRepository() FeatureRepository
	AudienceRepository() AudienceRepository
}
