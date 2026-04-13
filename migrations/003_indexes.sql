-- 003_indexes.sql
CREATE INDEX IF NOT EXISTS idx_items_user_id    ON items(user_id);
CREATE INDEX IF NOT EXISTS idx_items_created_at ON items(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_notif_user_read  ON notifications(user_id, read);
CREATE INDEX IF NOT EXISTS idx_notif_created    ON notifications(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_user       ON audit_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_audit_created    ON audit_logs(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_users_email      ON users(email);
-- Full-text search
ALTER TABLE items ADD COLUMN IF NOT EXISTS search_vec tsvector
  GENERATED ALWAYS AS (to_tsvector('english', name || ' ' || COALESCE(description,''))) STORED;
CREATE INDEX IF NOT EXISTS idx_items_fts ON items USING GIN(search_vec);
