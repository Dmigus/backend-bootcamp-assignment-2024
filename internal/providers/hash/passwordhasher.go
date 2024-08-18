package hash

import "golang.org/x/crypto/bcrypt"

type BCryptHasher struct{}

func (p *BCryptHasher) Hash(password string, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p *BCryptHasher) CheckPasswordHash(salt, password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	return err == nil
}
