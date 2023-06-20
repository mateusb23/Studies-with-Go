package main

import "fmt"

func main() {

	x := "Olá,"
	y := "bom dia!"

	fmt.Println(x, y)

	/* Diferente do fmt.Println , o fmt.Sprint ---> PODE SER UTILIZADO COMO VARIÁVEL <---
	conforme podemos ver nas linhas abaixo */

	z := fmt.Sprint(x, y)
	fmt.Println(z)

	z = fmt.Sprint(x, " ", y)
	fmt.Println(z)
}
