package hash

import "golang.org/x/crypto/bcrypt"

// Password hashes a plain-text password using bcrypt cost 12.
func Password(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 12)
	return string(bytes), err
}

// Verify returns true if the plain password matches the bcrypt hash.
func Verify(plain, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
