package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// structure for an API product with json tags
// swagger:model
type Product struct {
	// id for the product
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// name for the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// description for the product
	//
	// required: false
	// max length: 1000
	Description string `json:"description"`

	// price for the product
	//
	// required: true
	// min: 0
	Price float32 `json:"price" validate:"gt=0"`

	// sku (stock keeping unit) for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	// when the product was created
	//
	// required: false
	CreatedOn time.Time `json:"-"`

	// when the product was last updated
	//
	// required: false
	UpdatedOn time.Time `json:"-"`

	// when the product was deleted
	//
	// required: false
	DeletedOn time.Time `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-abcd-abcde

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1

}

type Products []*Product

// p*Products - method receiver. ToJSON method is attached to the Products type.
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// slice of pointers to Product structs
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
}
