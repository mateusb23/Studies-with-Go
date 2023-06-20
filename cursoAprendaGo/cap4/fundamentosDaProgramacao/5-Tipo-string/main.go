package main

import "fmt"

func main() {

	t := `Coffee
                  and
                tea,
                  please`
	fmt.Println(t)

	s1 := "\nHello, friends"
	fmt.Printf("%v ---> %T\n\n", s1, s1)

	s2 := s1 + ", coffee with milk?\n"
	fmt.Println(s2)

	sb := []byte(s1)
	fmt.Printf("%v ---> %T\n\n", sb, sb)

	/* OBS: strings em Go são IMUTÁVEIS, isso significa que não conseguimos pegar uma
	string a alterá-la, nós teríamos que criar uma outra string que contenha a string
	anterior com a alteração desejada, conforme podemos ver na declaração e atribuição
	da variável s2 na linha 16 acima */

}
