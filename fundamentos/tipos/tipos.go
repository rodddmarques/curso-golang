package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numericos interios
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal de 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "olá meu nome é Rodrigo"
	fmt.Println(s1 + "!!")
	fmt.Println("o tamanho da string s1 é", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Rodrigo`
	fmt.Println("o tamanho da string s2 é", len(s2))

}
