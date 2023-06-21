package repository

var products []model
var idNumber = 1

func NewProduct() Product {
	return &product{}
}

func (p *product) Add(name string) error {
	targetProduct := model{ID: idNumber, Name: name}
	products = append(products, targetProduct)
	idNumber++
	return nil
}

func (p *product) Remove(id int) error {
	targetId := -1
	for i := 0; i < len(products); i++ {
		targetProduct := products[i]
		if targetProduct.ID == id {
			targetId = i
			break
		}
	}
	if targetId != -1 {
		products = removeSlice(products, targetId)
	}
	return nil
}

func (p *product) GetAll() ([]model, error) {
	return products, nil
}

func removeSlice(slice []model, s int) []model {
	return append(slice[:s], slice[s+1:]...)
}
