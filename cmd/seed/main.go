package main

import ("database/sql"; "fmt"; "log"; "os"; "time")
_ "github.com/lib/pq"

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil { log.Fatal(err) }
	defer db.Close()
	// Seed admin user
	_, _ = db.Exec(`INSERT INTO users(email,password_hash,name,role,created_at) VALUES($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`
		, "admin@example.com", "$2a$10$placeholder", "Admin User", "admin", time.Now())
	// Seed sample items
	for i := 1; i <= 20; i++ {
		_, _ = db.Exec(`INSERT INTO items(name,description,user_id,created_at) VALUES($1,$2,1,$3) ON CONFLICT DO NOTHING`
			, fmt.Sprintf("Item %d", i), fmt.Sprintf("Sample item number %d", i), time.Now())
	}
	// Seed notifications
	_, _ = db.Exec(`INSERT INTO notifications(user_id,title,body,type,created_at) VALUES(1,$1,$2,$3,$4) ON CONFLICT DO NOTHING`
		, "Welcome!", "Your account is ready.", "info", time.Now())
	fmt.Println("seeded successfully")
}
