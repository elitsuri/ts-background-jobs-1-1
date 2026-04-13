#!/usr/bin/env bash
set -euo pipefail
echo '🚀 Setting up ts-background-jobs-1...'
command -v go >/dev/null || { echo 'Go not found. Install from https://go.dev'; exit 1; }
command -v psql >/dev/null || echo 'Warning: psql not found — DB setup may fail'
[ -f .env ] || cp .env.example .env && echo '✓ .env created'
source .env 2>/dev/null || true
go mod tidy && echo '✓ Go modules'
go build ./... && echo '✓ Build OK'
if command -v psql >/dev/null && [ -n "${DATABASE_URL:-}" ]; then
  go run ./cmd/migrate && echo '✓ Migrations applied'
  go run ./cmd/seed   && echo '✓ DB seeded'
else
  echo '⚠ Skipping DB setup — set DATABASE_URL in .env'
fi
echo '✅ Setup complete. Run: make dev'
