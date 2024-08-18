package salt

import "crypto/rand"

const saltSize = 16

type Generator struct{}

func (g *Generator) NewSalt() string {
	var salt = make([]byte, saltSize)
	_, _ = rand.Read(salt[:])
	return string(salt)
}
