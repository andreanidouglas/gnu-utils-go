package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	args := []string{"."}
	result := []string{}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	regex := "[a-zA-Z]+"

	for _, folder := range args {

		err := find(folder, regex, &result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}

}

func find(path string, nameCRegex string, result *[]string) []error {

	errSlice := []error{}

	fi, err := os.Stat(path)
	if err != nil {
		errSlice = append(errSlice, err)
		return errSlice
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		errSlice = append(errSlice, err)
		return errSlice
	}

	for _, file := range files {
		re, err := regexp.MatchString(nameCRegex, fi.Name())
		if err != nil {
			errSlice = append(errSlice, err)
			return errSlice
		}
		if file.IsDir() {
			//fmt.Println(path + "/" + file.Name())
			errS := find(path+"/"+file.Name(), nameCRegex, result)
			if err != nil {
				for _, err := range errS {
					errSlice = append(errSlice, err)
				}
			}
		}
		if re {
			*result = append(*result, path+"/"+file.Name())
		}

		for _, file := range *result {
			fmt.Printf("%s\n", file)
		}
	}
	if len(errSlice) > 0 {
		return errSlice
	}
	return nil
}
