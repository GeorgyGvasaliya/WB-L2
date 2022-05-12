package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	bufSize = 1024 * 8
)

// rewrite lib: https://github.com/hahnicity/go-wget
func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("no url")
	}
	path := flag.Arg(0)
	err := Wget(path, "file.html")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Wget(url, fileName string) error {
	resp, err := getResponse(url)
	if err != nil {
		return err
	}

	if fileName == "" {
		return errors.New("write filename")
	}

	err = writeToFile(fileName, resp)
	if err != nil {
		return err
	}

	return nil
}

// Make the GET request to a url, return the response
func getResponse(url string) (*http.Response, error) {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Write the response of the GET request to file
func writeToFile(fileName string, resp *http.Response) error {
	// Credit for this implementation should go to github user billnapier
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	if err != nil {
		return err
	}

	_, err = io.Copy(bufferedWriter, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
