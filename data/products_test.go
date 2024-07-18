package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1,
		SKU:   "nrw-epj-dro",
	}

	v := NewValidation()
	err := v.Validate(p)

	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name: "NED",
		SKU:  "nrw-epj-dro",
	}

	v := NewValidation()
	err := v.Validate(p)

	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "NED",
		Price: 1,
		SKU:   "nrw",
	}

	v := NewValidation()
	err := v.Validate(p)

	assert.Len(t, err, 1)
}

func TestValidProductDoesNotReturnErr(t *testing.T) {
	p := Product{
		Name:  "NED",
		Price: 1,
		SKU:   "nrw-epj-dro",
	}

	v := NewValidation()
	err := v.Validate(p)

	assert.Len(t, err, 1)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		&Product{
			ID:          1,
			Name:        "NED",
			Description: "My favorite",
			Price:       1,
			SKU:         "nrw-epj-dro",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}
