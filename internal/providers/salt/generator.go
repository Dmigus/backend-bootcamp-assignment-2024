package salt

import "crypto/rand"

const saltSize = 16

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) NewSalt() []byte {
	var salt = make([]byte, saltSize)
	_, _ = rand.Read(salt[:])
	return salt
}
