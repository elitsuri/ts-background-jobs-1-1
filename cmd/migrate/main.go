package main

import ("database/sql"; "fmt"; "log"; "os"; "path/filepath"; "sort"; "strings")
_ "github.com/lib/pq"

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" { log.Fatal("DATABASE_URL not set") }
	db, err := sql.Open("postgres", dsn)
	if err != nil { log.Fatal(err) }
	defer db.Close()
	if err := db.Ping(); err != nil { log.Fatal("DB ping failed:", err) }
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (version TEXT PRIMARY KEY, applied_at TIMESTAMPTZ DEFAULT NOW())`)
	files, _ := filepath.Glob("migrations/*.sql")
	sort.Strings(files)
	for _, f := range files {
		ver := strings.TrimSuffix(filepath.Base(f), ".sql")
		var exists bool
		_ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version=$1)", ver).Scan(&exists)
		if exists { fmt.Println("skip:", ver); continue }
		content, err := os.ReadFile(f)
		if err != nil { log.Fatal("read:", err) }
		if _, err := db.Exec(string(content)); err != nil { log.Fatalf("migrate %s: %v", ver, err) }
		_, _ = db.Exec("INSERT INTO schema_migrations(version) VALUES($1)", ver)
		fmt.Println("applied:", ver)
	}
	fmt.Println("migrations complete")
}
