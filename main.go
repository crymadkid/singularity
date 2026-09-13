package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "использование: sgl <длина>")
		os.Exit(1)
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length < 1 {
		fmt.Fprintln(os.Stderr, "длина должна быть положительным числом")
		os.Exit(1)
	}

	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			fmt.Fprintln(os.Stderr, "не удалось сгенерировать пароль:", err)
			os.Exit(1)
		}
		password[i] = symbols[index.Int64()]
	}

	fmt.Println(string(password))
}
