package feature

const LIST_SQL = `
SELECT features.name,
	   features.label,
	   features.enabled,
	   features.has_audience,
	   features.has_audience_group,
	   features.created_at,
	   features.updated_at,
	   features.enabled_at
FROM features
`

const GET_BY_SQL = `
SELECT features.name,
	   features.label,
	   features.enabled,
	   features.has_audience,
	   features.has_audience_group,
	   features.created_at,
	   features.updated_at,
	   features.enabled_at
FROM features
`

const SAVE_SQL = `
INSERT INTO features (name, label, enabled,
			has_audience, has_audience_group, created_at,
			updated_at, enabled_at)
VALUES (:name, :label, :enabled,
	:has_audience, :has_audience_group, :created_at,
	:updated_at, :enabled_at)
ON CONFLICT (name)
DO
   UPDATE SET name = :name, label = :label, enabled = :enabled,
		has_audience = :has_audience, has_audience_group = :has_audience_group, created_at = :created_at,
		updated_at = :updated_at, enabled_at = :enabled_at;
`

const SIZE_SQL = `
SELECT COUNT(*) FROM features
`
const GET_SQL = `
SELECT features.name,
	   features.label,
	   features.enabled,
	   features.has_audience,
	   features.has_audience_group,
	   features.created_at,
	   features.updated_at,
	   features.enabled_at
FROM features
WHERE name = $1;
`

const DELETE_SQL = `
DELETE FROM features
WHERE name = $1;
`
