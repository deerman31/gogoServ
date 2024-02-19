package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func pathSet() (string, error) {
	size := len(os.Args)
	if size > 2 {
		return "", errors.New("ERROR: 引数は1つもしくは2つです!")
	} else if size == 1 {
		return "conf/default.conf", nil
	}
	if (os.Args[1] == ".conf") {
		return "", errors.New("ERROR: 拡張子は.confです!")
	}
	index := strings.LastIndex(os.Args[1], ".")
	if index == -1 {
		return "", errors.New("ERROR: 拡張子は.confです!")
	}
	const exe string = ".conf"
	s := os.Args[1][index:]
	if len(s) != len(exe) {
		return "", errors.New("ERROR: 拡張子は.confです!")
	}
	for i := 0; i < len(exe) && i < len(s); i += 1 {
		if exe[i] != s[i] {
			return "", errors.New("ERROR: 拡張子は.confです!")
		}
	}

	return os.Args[1], nil
}

func fileOpen(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	path, err := pathSet()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	text, err := fileOpen(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(text)
}
