package db

type User struct {
	id       uint
	name     string
	password string
}

type Product struct {
	id          uint
	name        string
	description string
	cost        float32
	thumbnail   string
}

type Cart struct {
	id       uint
	products []Product
	cost     float32
}
