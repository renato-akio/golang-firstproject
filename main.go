package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/renato-akio/golang-firstproject/internal/infra/database"
	"github.com/renato-akio/golang-firstproject/internal/usecase"
)

// comando para executar o arquivo: go run main.go
// comando para gerar o build: go build main.go
// comando para inicializar o modulo do projeto: go mod init github.com/renato-akio/golang-firstproject
// gerar build para windows  GOOS=windows go build main.go
// atualizar pacote: go mod tidy
// executar sqlite3: sqlite3 db.sqlite3
// _ "github.com/mattn/go-sqlite3"  quando nao está utilizando o pacote diretamente, para nao ficar gerando erro
// utilizar o _ na importacao do pacote

type Car struct {
	Model string
	Color string
}

// metodo
// (c Car) => cria uma associacao com Car
func (c Car) Start() {
	println(c.Model + " has been started")
}

// (c *Car) => Utilizando o *  significa que irá modificar o valor do objeto na memoria
// (c Car) => Utilizando o *  significa que não irá modificar o valor do objeto na memoria, mas criar uma cópia
// e o valor se mantido, até o final da execução do método
func (c Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

func (c *Car) ChangeColorMemory(color string) {
	c.Color = color
	println("New color memory: " + c.Color)
}

// funcao
func soma(x, y int) int {
	return x + y
}

func main() {
	/*
		println("Hello, world!")

		car := Car{
			Model: "Ferrari",
			Color: "Red",
		}
		car.Model = "Fiat"
		println(car.Model)

		car.Start()
		car.ChangeColor("Blue")

		println(car.Color)

		car.ChangeColorMemory("Green")
		println(car.Color)

		a := 10
		b := a // b é uma copia do valor de a, mas criou o seu proprio espaco de memoria
		b = 20

		println(&a) // retorno o valor do endereço da memória
		println(a)
		println(b)
		println(a)

		order, err := entity.NewOrder("1", -10, 1)

		if err != nil {
			println(err.Error())
		} else {
			println(order.ID)
		}
	*/

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
	fmt.Println(output)
}
