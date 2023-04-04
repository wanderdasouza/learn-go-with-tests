package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

	if sum != expected {
		// Para printar os números como inteiro ao invés de string,
		// use o %d
		t.Errorf("want '%d', but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// O comentário Output abaixo é obrigatório para que o exemplo seja executado de
	// forma correta e exibido na documentação

	// Output: 6
}