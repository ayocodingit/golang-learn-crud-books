package book

type Service interface {
	Index() ([]Book, error)
	Store(bookRequest BookRequest) (Book, error)
	Show(ID int) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Destroy(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Book, error) {
	books, err := s.repository.Index()
	return books, err
}

func (s *service) Store(bookRequest BookRequest) (Book, error) {
	price, err := bookRequest.Price.Int64()
	rating, err := bookRequest.Rating.Int64()
	discount, err := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Store(book)

	return newBook, err
}

func (s *service) Show(ID int) (Book, error) {
	book, err := s.repository.Show(ID)
	return book, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	price, err := bookRequest.Price.Int64()
	rating, err := bookRequest.Rating.Int64()
	discount, err := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	updateBook, err := s.repository.Update(book)

	return updateBook, err
}

func (s *service) Destroy(ID int) (Book, error) {
	book, err := s.repository.Show(ID)
	deleteBook, err := s.repository.Destroy(book)
	return deleteBook, err
}
