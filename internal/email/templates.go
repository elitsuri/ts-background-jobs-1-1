package email

const WelcomeTemplate = `<!DOCTYPE html><html><body>
<h1>Welcome to {{.AppName}}!</h1>
<p>Hello {{.Name}}, your account is ready.</p>
<a href='{{.URL}}'>Get started</a>
</body></html>`

const PasswordResetTemplate = `<!DOCTYPE html><html><body>
<h1>Password Reset</h1>
<p>Click below to reset your password (expires in 1 hour):</p>
<a href='{{.ResetURL}}'>Reset Password</a>
</body></html>`

const NotificationTemplate = `<!DOCTYPE html><html><body>
<h2>{{.Title}}</h2>
<p>{{.Body}}</p>
</body></html>`
