package application

type ProductService struct {
	Persistence IProductPersistent
}

func (s *ProductService) Get(id string) (IProduct, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
