CREATE TABLE IF NOT EXISTS contents (
  id           SERIAL PRIMARY KEY,
  title        VARCHAR(200) NOT NULL,
  excerpt      VARCHAR(200) UNIQUE NOT NULL,
  description  TEXT NOT NULL,
  image        TEXT NULL,
  status       VARCHAR(20) NOT NULL DEFAULT 'PUBLISH',
  tags         TEXT NOT NULL,
  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  created_by   INT REFERENCES users(id) ON DELETE CASCADE,
  updated_by   INT REFERENCES users(id) ON DELETE CASCADE,
  category_id  INT REFERENCES categories(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_contents_created_by_id ON contents(created_by);
CREATE INDEX IF NOT EXISTS idx_contents_category_id   ON contents(category_id);