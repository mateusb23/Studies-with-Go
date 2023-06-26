package main

import "fmt"

func main() {
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}
	fatia := sabores[0:2]
	fmt.Println(fatia)

	dias := []string{"domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"}
	fatiadias := dias[2:len(dias)]
	fatiadias1 := dias[:2]
	fatiadias2 := dias[2:]
	fmt.Println(fatiadias)
	fmt.Println(fatiadias1)
	fmt.Println(fatiadias2)

	fmt.Println("-----------------------------------------------")

	// COMO ACESSAR TODOS OS ÍNDICES DE UM slice SEM PRECISAR UTILIZAR O range:
	todososdias := dias[:]
	fmt.Println(todososdias)

	// OU
	for i := 0; i < len(dias); i++ {
		fmt.Println(dias[i])
	}

	// OU
	for _, v := range dias {
		fmt.Println(v)
	}

	// COMO DELETAR VALORES DE UM slice, (nesse caso deletei o valor "abacaxi"):
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

}
