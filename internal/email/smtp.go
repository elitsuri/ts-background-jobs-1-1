package email

import ("fmt"; "net/smtp"; "os"; "strings")

type SMTPClient struct{ host, port, user, pass, from string }

func NewSMTPClient() *SMTPClient {
	return &SMTPClient{
		host: os.Getenv("SMTP_HOST"), port: os.Getenv("SMTP_PORT"),
		user: os.Getenv("SMTP_USER"), pass: os.Getenv("SMTP_PASS"),
		from: os.Getenv("SMTP_FROM"),
	}
}

func (c *SMTPClient) Send(to, subject, htmlBody string) error {
	if c.host == "" { return nil }
	headers := strings.Join([]string{
		"From: " + c.from, "To: " + to, "Subject: " + subject,
		"MIME-Version: 1.0", "Content-Type: text/html; charset=UTF-8", "",
	}, "\r\n")
	auth := smtp.PlainAuth("", c.user, c.pass, c.host)
	return smtp.SendMail(fmt.Sprintf("%s:%s", c.host, c.port), auth, c.from, []string{to}, []byte(headers+htmlBody))
}
