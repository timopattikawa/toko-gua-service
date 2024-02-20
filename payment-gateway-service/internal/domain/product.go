package domain

type Product struct {
	Id    int    `db:"id"  json:"id"`
	Name  string `db:"name" json:"name"`
	Price int    `db:"price" json:"price"`
}
