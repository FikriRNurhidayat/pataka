CREATE TABLE IF NOT EXISTS feature_audiences (
  feature_name       VARCHAR(255) NOT NULL,
  audience_id        VARCHAR(255) NOT NULL,
  enabled            BOOLEAN      NOT NULL,
  created_at         TIMESTAMP    NOT NULL,
  updated_at         TIMESTAMP    NOT NULL,
  enabled_at         TIMESTAMP,
  PRIMARY KEY (feature_name, audience_id)
);
