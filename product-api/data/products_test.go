package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1,
		SKU:   "abc-bca-cbaabc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
