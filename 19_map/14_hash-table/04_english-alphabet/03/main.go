package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[int]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	numberWord := 0
	for sc.Scan() {
		words[numberWord] = sc.Text()
		numberWord++
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	var keys []int
	for key := range words {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	i := 0
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", words[k])
		if i == 200{
			break
		}
		i++
	}

}

