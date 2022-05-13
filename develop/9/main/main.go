package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	//формируем запрос
	ctx, cancel_func := context.WithTimeout(context.Background(), timeout)
	request = request.WithContext(ctx)

	//выполнение запроса
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	//закрываем соединение
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		cancel_func()
		return nil, fmt.Errorf("INVALID RESPONSE; status: %s", response.Status)
	}

	return ioutil.ReadAll(response.Body)
}

func main() {
	//флаги со значениями по умолчинию
	url := flag.String("url", "https://habr.com/", "url")
	timeout := flag.Duration("timeout", 5*time.Second, "string1")
	output_path := flag.String("output", "file.txt", "string2")
	flag.Parse()

	//получаем данные
	content, err := HTTPGet(*url, *timeout*time.Second)
	if err != nil {
		log.Fatalln("HTTPGET: ", err)
	}
	//запись в файл, с разрешениями
	err = ioutil.WriteFile(*output_path, content, 0644)
	if err != nil {
		log.Fatalln("WriteFile: ", err)
	}
	return
}
