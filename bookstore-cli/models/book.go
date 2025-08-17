package models


type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Image_url string `json:"imageUrl"`
}