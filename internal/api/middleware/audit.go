package middleware

import ("database/sql"; "net/http"; "strings")

var auditDB *sql.DB
func SetAuditDB(db *sql.DB) { auditDB = db }

func Audit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutating := map[string]bool{"POST":true,"PUT":true,"PATCH":true,"DELETE":true}
		if mutating[r.Method] && auditDB != nil {
			userID, _ := r.Context().Value(ctxKeyUserID).(int64)
			parts := strings.Split(r.URL.Path, "/")
			resource := ""
			if len(parts) > 3 { resource = parts[3] }
			_, _ = auditDB.Exec(
				`INSERT INTO audit_logs(user_id,action,resource_type,ip_address) VALUES($1,$2,$3,$4)`,
				userID, r.Method, resource, r.RemoteAddr,
			)
		}
		next.ServeHTTP(w, r)
	})
}
