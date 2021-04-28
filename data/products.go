package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func validateSKU(fl validator.FieldLevel) bool {
	//sku is of format abc-edsd-fkdjs
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	// in-case no match is found
	if len(matches) != 1 {
		return false
	}
	return true
}

func (p *Product) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productLists
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productLists = append(productLists, p)
}

func getNextID() int {
	lp := productLists[len(productLists)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productLists[pos] = p
	return nil
}

var ErrorProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (int, error) {
	for i, p := range productLists {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrorProductNotFound
}

var productLists = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
