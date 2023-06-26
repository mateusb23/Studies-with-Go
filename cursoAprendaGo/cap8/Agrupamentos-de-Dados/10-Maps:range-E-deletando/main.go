package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}
	fmt.Println(qualquercoisa)

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}

	fmt.Println("Total =", total)

	// DELETAR DE UM map
	delete(qualquercoisa, 983)
	fmt.Println(qualquercoisa)
}
