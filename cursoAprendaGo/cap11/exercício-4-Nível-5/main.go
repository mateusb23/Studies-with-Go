package main

import "fmt"

func main() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Suco de Laranja",
		sabor:   "doce",
		ondetem: []string{"EUA", "Brasil", "Colômbia", "Uruguai", "Angola"},
		vaibemcom: map[string]string{
			"no café da manhã": "ovos mexidos",
			"no almoço":        "peixe assado",
			"no jantar":        "panquequas",
		},
	}

	fmt.Println(x)
}
