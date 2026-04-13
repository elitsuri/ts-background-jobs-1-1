# ts-background-jobs-1 - API Reference

## Authentication
All endpoints except `/auth/*` require `Authorization: Bearer <token>` header.

## Endpoints

### Auth
| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v1/auth/register` | Register new user |
| POST | `/api/v1/auth/login` | Login, returns JWT tokens |
| POST | `/api/v1/auth/refresh` | Refresh access token |

### Items
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/items` | OK | List items (paginated) |
| POST | `/api/v1/items` | OK | Create item |
| GET | `/api/v1/items/{id}` | OK | Get item by ID |
| PUT | `/api/v1/items/{id}` | OK | Update item |
| DELETE | `/api/v1/items/{id}` | OK | Delete item |

### Analytics
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/analytics/overview` | OK | Total counts |
| GET | `/api/v1/analytics/timeseries` | OK | Daily activity (param: days) |

### Notifications
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/notifications` | OK | List notifications |
| POST | `/api/v1/notifications` | OK | Create notification |
| PUT | `/api/v1/notifications/{id}/read` | OK | Mark as read |

### Admin (admin role required)
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/admin/users` | admin | List all users |
| PUT | `/api/v1/admin/users/{id}/role` | admin | Update user role |
| DELETE | `/api/v1/admin/users/{id}` | admin | Delete user |

### Search
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/search?q=&type=` | OK | Search items/users |

### Upload
| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/upload/file` | OK | Upload file (multipart, max 10MB) |

## WebSocket
Connect to `ws://host/ws` with Authorization header. Receives JSON events:

```json
{"type": "notification", "title": "...", "body": "..."}
```
