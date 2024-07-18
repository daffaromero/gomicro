package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "NED",
		Price: 1,
		SKU:   "nrw-epj-dro",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
