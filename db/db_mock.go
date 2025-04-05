package db

import "slices"

var products []Product

func InitMockProducts() {
	products = append(products, Product{
		id:          1,
		name:        "Toy car",
		description: "A simple car shaped toy",
		thumbnail:   "🚗",
		cost:        10.20,
	})
	products = append(products, Product{
		id:          2,
		name:        "Apples",
		description: "How about them apples",
		thumbnail:   "🍎",
		cost:        2.04,
	})
	products = append(products, Product{
		id:          3,
		name:        "Cookies",
		description: "Hand over those cookies",
		thumbnail:   "🍪",
		cost:        4.23,
	})
}

func GetProducts() []Product {
	return slices.Clone(products)
}

func GetProduct(id uint) Product {
	index := slices.IndexFunc(products, func(product Product) bool { return product.id == id })
	return products[index]
}

var cart = Cart{
	id:       1,
	products: []Product{products[1]},
	cost:     products[1].cost,
}

func GetCart() Cart {
	return cart
}

func calCartCost() {
	cart.cost = 0
	for _, p := range cart.products {
		cart.cost += p.cost
	}
}

func AddToCart(id uint) {
	product := GetProduct(id)
	cart.products = append(cart.products, product)
	calCartCost()
}

func RemoveFromCart(id uint) {
	cart.products = slices.Delete(cart.products, int(id)-1, int(id))
	calCartCost()
}
