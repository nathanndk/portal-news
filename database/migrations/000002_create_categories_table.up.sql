CREATE TABLE IF NOT EXISTS categories (
  id          SERIAL PRIMARY KEY,
  title       VARCHAR(200) NOT NULL,
  slug        VARCHAR(200) UNIQUE NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP NULL  DEFAULT CURRENT_TIMESTAMP,
  created_by  INT REFERENCES users(id) ON DELETE CASCADE,
  updated_by  INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_categories_created_by_id ON categories(created_by);
