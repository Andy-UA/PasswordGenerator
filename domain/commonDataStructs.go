package domain

type PasswordConfig struct {
	MinLength          int
	SpecialCharsAmount int
	NumberAmount       int
}

type Response struct {
	GeneratedPassword string
}
