package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	err := cp(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

}

func cp(src, dest string) error {

	stat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("could not stat file")
	}

	if stat.Mode().IsDir() {
		return fmt.Errorf("omitting directory %s", src)
	}

	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("cannot open file %s", src)
	}

	newFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("cannot create file %s", dest)
	}

	re := bufio.NewReader(file)
	wr := bufio.NewWriter(newFile)

	bytesCount, err := io.Copy(wr, re)
	if err != nil {
		return fmt.Errorf("cannot create file %s", dest)
	}
	_ = bytesCount
	return nil
}
