package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"bytes"
	"math/rand"
	"time"
	"errors"
	"unicode"
)

func main() {
	var err error

	nfaces, err := readFacesInput()
	if (err != nil) {
		fmt.Printf("%q\n", err)
		os.Exit(1)	
	}

	rerollChoice := true

	for rerollChoice {
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("Rolling a %v-face dice\n", nfaces)
		fmt.Printf("-----\nResult: %v\n-----\n", rand.Intn(nfaces))
		fmt.Printf("Wanna reroll a %v-face dice? [Y] or [n]\n", nfaces)
		rerollChoice, err = readRerollChoice()
		if (err != nil) {
			fmt.Printf("%q\n", err)
			os.Exit(1)	
		}
	}

}

func readFacesInput() (int, error) {
	var err error
	err = nil

	reader := io.Reader(os.Stdin)
	input := make([]byte, 5)
	fmt.Printf("How many faces? ")
	_, err = reader.Read(input)
	if (err != nil) {
		return 0, err
	}
	cleanInput := strings.Replace(string(bytes.Trim(input, "\x00")), "\r\n", "", -1)

	nfaces, err := strconv.Atoi(cleanInput)
	if (err != nil) {
		return 0, err
	}
	return nfaces, err
}

func readRerollChoice() (bool, error){
	var err error
	err = nil

	var rerollChoice bool

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if (err != nil) {
		return false, err
	}
	char = unicode.ToLower(char)

	switch char {
	case 'y':
	  rerollChoice = true
	  break
	case 'n':
	  rerollChoice = false
	  break
	case 13:
		rerollChoice = true
		break
	default:
		err = errors.New("Wrong entry")
	}
	return rerollChoice, err
}