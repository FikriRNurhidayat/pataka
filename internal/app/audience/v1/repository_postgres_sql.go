package audience

const SIZE_SQL = `
SELECT COUNT(*)
FROM feature_audiences
`

const DELETE_SQL = `
DELETE FROM feature_audiences
WHERE feature_name = $1
AND audience_id = $2;
`

const GET_SQL = `
SELECT feature_audiences.feature_name,
	   feature_audiences.audience_id,
	   feature_audiences.enabled,
	   feature_audiences.created_at,
	   feature_audiences.updated_at,
	   feature_audiences.enabled_at
FROM feature_audiences
WHERE feature_name = $1
AND audience_id = $2;
`

const LIST_SQL = `
SELECT feature_audiences.feature_name,
	   feature_audiences.audience_id,
	   feature_audiences.enabled,
	   feature_audiences.created_at,
	   feature_audiences.updated_at,
	   feature_audiences.enabled_at
FROM feature_audiences
`

const SAVE_SQL = `
INSERT INTO feature_audiences (
	feature_name,
	audience_id,
	enabled,
	created_at,
	updated_at,
	enabled_at
)
VALUES (
	:feature_name,
	:audience_id,
	:enabled,
	:created_at,
	:updated_at,
	:enabled_at
)
ON CONFLICT (feature_name, audience_id)
DO
   UPDATE SET feature_name = :feature_name,
			  audience_id = :audience_id,
			  enabled = :enabled,
			  created_at = :created_at,
			  updated_at = :updated_at,
			  enabled_at = :enabled_at;
`
