package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	accepts := flag.String("a", "", "Accepted words.")
	flag.Parse()

	arg := flag.Arg(0)
	if arg == "" {
		log.Fatalln("Need 1 arg.")
	}
	digits, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatalln("Digit must be int.")
	}

	mrand.Seed(time.Now().UnixNano())

	key := GenerateRandomKey(digits)

	result := ""
	for _, v := range key {
		result += string(adjustString(v, *accepts))
	}
	fmt.Println(result)
}

func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}

	return k
}

func adjustString(n byte, accepts string) byte {
	switch {
	case 48 <= n && n <= 57:
		// 1-9
		return n
	case 65 <= n && n <= 90:
		// A-Z
		return n
	case 97 <= n && n <= 112:
		// a-z
		return n
	case strings.Contains(accepts, string(n)):
		return n
	}
	return adjustString(byte((int(n)+mrand.Intn(9))*mrand.Intn(9)%112), accepts)
}
