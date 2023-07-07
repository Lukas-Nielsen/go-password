package password

import "golang.org/x/crypto/bcrypt"

// Verifies that a password matches a hash
func Verify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
