package hash

import "golang.org/x/crypto/bcrypt"

type BCryptHasher struct{}

func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{}
}

func (p *BCryptHasher) Hash(salt []byte, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(p.saltWithPassword(salt, password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p *BCryptHasher) CheckPasswordHash(salt []byte, password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), p.saltWithPassword(salt, password))
	return err == nil
}

func (p *BCryptHasher) saltWithPassword(salt []byte, password string) []byte {
	return append(salt, password...)
}
