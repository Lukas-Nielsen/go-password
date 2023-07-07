package password

import "golang.org/x/crypto/bcrypt"

// creates a new password hash using a strong one-way hashing algorithm
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
