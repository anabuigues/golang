package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sort"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	numberWord := 0
	for sc.Scan() {
		words[sc.Text()] = strconv.Itoa(numberWord) // get token
		numberWord++
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, value := range words {
		fmt.Println(k, value)
		if i == 200 {
			break
		}
		i++
	}

}
