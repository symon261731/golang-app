package link

import "math/rand"

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: generateHashLink(6),
	}
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPSADFGHJKLZXZCVBNM")

func generateHashLink(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
