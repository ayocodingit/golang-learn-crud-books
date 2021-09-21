package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
}

func NewResponse(book Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		Description: book.Description,
		Rating:      book.Rating,
		Discount:    book.Discount,
	}
}
