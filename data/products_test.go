package data

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 2)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: -1,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 2)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: 1.22,
		SKU:   "test",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNotReturnsErr(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Nil(t, err)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		&Product{
			Name: "Test",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}
