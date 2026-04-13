# ts-background-jobs-1 - Architecture

## Layer Diagram
```

           Web / SPA (web/static/)

     API Handlers (internal/api/handlers/)

    Services (internal/service/)

Repositories (internal/repository/)

       PostgreSQL Database

```

## Package Structure
```
cmd/
server/     - HTTP server entry point
migrate/    - DB migration binary
seed/       - DB seeder binary
internal/
api/        - HTTP layer (router, handlers, middleware)
config/     - Configuration from env
database/   - DB connection + embedded migrations
domain/     - Core domain types
model/      - DB row types
repository/ - DB access layer
service/    - Business logic
worker/     - Background workers
events/     - Pub/sub event bus
cache/      - In-memory + Redis cache
websocket/  - WS hub + client
pkg/
jwt/        - JWT helpers
hash/       - Password hashing
response/   - JSON response helpers
migrations/   - SQL migration files
web/static/   - SPA dashboard
```

## Tech Stack
- **Language**: Go 1.22+
- **HTTP**: stdlib net/http with pattern routing
- **Database**: PostgreSQL via lib/pq
- **Auth**: JWT (golang-jwt/jwt/v5)
- **Password**: bcrypt (golang.org/x/crypto)
- **Hot reload**: Air
- **Migrations**: embedded SQL via go:embed
