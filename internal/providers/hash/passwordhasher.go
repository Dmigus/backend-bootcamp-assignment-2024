package hash

import "golang.org/x/crypto/bcrypt"

type BCryptHasher struct{}

func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{}
}

func (p *BCryptHasher) Hash(salt []byte, password string) ([]byte, error) {
	salted := p.saltWithPassword(salt, password)
	return bcrypt.GenerateFromPassword(salted, bcrypt.DefaultCost)
}

func (p *BCryptHasher) CheckPasswordHash(salt []byte, password string, hash []byte) bool {
	salted := p.saltWithPassword(salt, password)
	err := bcrypt.CompareHashAndPassword(hash, salted)
	return err == nil
}

func (p *BCryptHasher) saltWithPassword(salt []byte, password string) []byte {
	return append(salt, password...)
}
