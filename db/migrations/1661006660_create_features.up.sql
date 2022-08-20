CREATE TABLE IF NOT EXISTS features (
  name               VARCHAR(255) NOT NULL,
  label              VARCHAR(255) NOT NULL,
  enabled            BOOLEAN      NOT NULL,
  has_audience       BOOLEAN      NOT NULL,
  has_audience_group BOOLEAN      NOT NULL,
  created_at         TIMESTAMP    NOT NULL,
  updated_at         TIMESTAMP    NOT NULL,
  enabled_at         TIMESTAMP,
  PRIMARY KEY (name)
);
