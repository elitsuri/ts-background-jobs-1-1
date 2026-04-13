package service

import ("crypto/tls"; "fmt"; "net/smtp"; "os"; "strings")

type EmailService struct{ host, port, user, pass, from string }

func NewEmailService() *EmailService {
	return &EmailService{
		host: os.Getenv("SMTP_HOST"),
		port: os.Getenv("SMTP_PORT"),
		user: os.Getenv("SMTP_USER"),
		pass: os.Getenv("SMTP_PASS"),
		from: os.Getenv("SMTP_FROM"),
	}
}

func (s *EmailService) Send(to, subject, body string) error {
	if s.host == "" { return nil } // not configured
	addr := s.host + ":" + s.port
	auth := smtp.PlainAuth("", s.user, s.pass, s.host)
	msg := strings.Join([]string{
		"From: " + s.from,
		"To: " + to,
		"Subject: " + subject,
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"",
		body,
	}, "\r\n")
	tlsConf := &tls.Config{InsecureSkipVerify: false, ServerName: s.host}
	_ = tlsConf
	return smtp.SendMail(addr, auth, s.from, []string{to}, []byte(msg))
}

func (s *EmailService) SendWelcome(to, name string) error {
	return s.Send(to, "Welcome!", fmt.Sprintf("<h1>Welcome, %s!</h1>", name))
}
