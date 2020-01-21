package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("Begini script Detail CEP.")
	fmt.Println("")
	time.Sleep(1000 * time.Millisecond)

	var CEP string
	fmt.Print("Enter your CEP: ")
	fmt.Scanln(&CEP)

	var url strings.Builder
	url.WriteString("https://viacep.com.br/ws/")
	url.WriteString(CEP)
	url.WriteString("/json/")
	fmt.Println("Execing GET of", url.String())
	fmt.Println("")

	response, err := http.Get(url.String())
	if err != nil { //Se tiver error na request.
		log.Fatalln(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil { //Se tiver error no body.
		log.Fatalln(err)
	}

	fmt.Println(string(body))
	time.Sleep(10000 * time.Millisecond)
}
