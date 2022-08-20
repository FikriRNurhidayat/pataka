package repository

const LIST_SQL = `
SELECT features.name,
	   features.label,
	   features.enabled,
	   features.created_at,
	   features.updated_at,
	   features.enabled_at
FROM features
`

const SAVE_SQL = `
INSERT INTO features (name, label, enabled, created_at, updated_at, enabled_at)
VALUES (:name, :label, :enabled, :created_at, :updated_at, :enabled_at)
ON CONFLICT (name)
DO
   UPDATE SET name = :name, label = :label, enabled = :enabled, created_at = :created_at, updated_at = :updated_at, enabled_at = :enabled_at;
`

const SIZE_SQL = "SELECT COUNT(*) FROM features"

const GET_SQL = `
SELECT features.name,
	   features.label,
	   features.enabled,
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
