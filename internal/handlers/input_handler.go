package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your choice: ")
	input, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.TrimSpace(input))
}

func GetTaskName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter task name: ")
	input, _ := reader.ReadString('\n')
	return ToTitle(strings.TrimSpace(input))
}

func GetDelayTime() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter delay time (in seconds): ")
	input, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.TrimSpace(input))
}

func ToTitle(input string) string {
	words := strings.Fields(strings.ToLower(input))
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}
