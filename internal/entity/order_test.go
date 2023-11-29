package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// para criar uma funcao de teste, obrigatoriamente o nome da funcao deve iniciar com Test
// para executar o teste: go test ./...
// para baixar um pacote externo : go mod tidy

/* teste com uma funcao de modo nativa do GO
func TestIfItGetsAnErrorIfIDISBlank(t *testing.T) {
	order := Order{}
	if order.Validate() == nil {
		t.Error("ID is required")
	}
}
*/

func TestIfItGetsAnErrorIfIDISBlank(t *testing.T) {
	order := Order{}
	err := order.Validate()
	assert.EqualError(t, err, "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	println("aqui")
	order := Order{ID: "123"}
	err := order.Validate()

	assert.EqualError(t, err, "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	println("aqui")
	order := Order{ID: "123", Price: 10.0}
	err := order.Validate()

	assert.EqualError(t, err, "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinaPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
