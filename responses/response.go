package responses

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/akecel/gabbro/utils"

	"github.com/martinlindhe/imgcat/lib"
)

func PrintImageResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	imageReader := bytes.NewReader(imageData)

	fmt.Println("")
	if err := imgcat.Cat(imageReader, os.Stdout); err != nil {
		panic(err)
	}
}

func PrintHeader(header string) {
	title := utils.SetColor()
	fmt.Println("\n" + title(header) + ":")
}

func PrintResponse(s interface{}) {
	title := utils.SetColor()

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Name
		value := v.Field(i).Interface()

		if value != "" {
			fmt.Printf(title("%s:")+" %v\n", utils.CamelCaseToNormal(field), value)
		}
	}
	fmt.Println("")
}

func PrintErrorResponse(message string, exit ...bool) {
	var haveToExit bool = true
	if len(exit) > 0 {
		haveToExit = exit[0]
	}

	fmt.Println("\n" + message)

	if haveToExit == true {
		os.Exit(0)
	}
}
