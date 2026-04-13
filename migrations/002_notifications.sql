-- 002_notifications.sql
CREATE TABLE IF NOT EXISTS notifications (
  id         BIGSERIAL PRIMARY KEY,
  user_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title      TEXT NOT NULL,
  body       TEXT NOT NULL DEFAULT '',
  type       TEXT NOT NULL DEFAULT 'info',
  read       BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS email_queue (
  id          BIGSERIAL PRIMARY KEY,
  to_address  TEXT NOT NULL,
  subject     TEXT NOT NULL,
  body        TEXT NOT NULL,
  attempts    INT NOT NULL DEFAULT 0,
  sent        BOOLEAN NOT NULL DEFAULT FALSE,
  sent_at     TIMESTAMPTZ,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS audit_logs (
  id            BIGSERIAL PRIMARY KEY,
  user_id       BIGINT,
  action        TEXT NOT NULL,
  resource_type TEXT,
  resource_id   TEXT,
  ip_address    TEXT,
  user_agent    TEXT,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
