package models

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Price     string `json:"price"`
	Image_url string `json:"imageUrl"`
}
