package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const apiKey = "apiKey"
const apiURL = "https://uork.org/search/status/check-account.php?apikey=%s&id="

func getUserInput(quest string) (string, error) {
	fmt.Print(quest)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	userInput, err := getUserInput("Insira o ID ou e-mail do usuário: ")
	if err != nil {
		fmt.Println("Erro ao obter entrada:", err)
		return
	}

	apiEndpoint := fmt.Sprintf(apiURL, apiKey) + userInput

	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("Erro na solicitação:", err)
		return
	}
	defer resp.Body.Close()

	var info string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		info += scanner.Text()
	}

	fmt.Println(info)
}
