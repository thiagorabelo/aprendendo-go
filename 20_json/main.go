package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` /* `json:"-"` - Faz a conversão ser ignorada */
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func marshal() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroBytes, err := json.Marshal(c) // []byte, error
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(cachorroBytes)
	fmt.Println(buffer)

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	fmt.Println(c2)
	cachorro2Bytes, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	buffer2 := bytes.NewBuffer(cachorro2Bytes)
	fmt.Println(buffer2)
}

func unmarshal() {
	cachorroEmJson := `{ "nome": "Rex", "raca": "Dálmata", "Idade": 3 }`
	fmt.Println(cachorroEmJson)

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroEmJson), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	cachorro2EmJson := `{ "nome": "Toby", "raca": "Poodle" }`
	fmt.Println(cachorro2EmJson)

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(cachorro2EmJson), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}

func main() {
	fmt.Println("# Marshal")
	marshal()

	fmt.Println("\n# Unmarshal")
	unmarshal()
}
