package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Nic",
		Price: 1.10,
		SKU:   "asd-adq-sas",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
