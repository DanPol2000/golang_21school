package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type recipes struct {
	Cake []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name" xml:"itemname"`
			IngredientCount string `json:"ingredient_count" xml:"itemcount"`
			IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

type JSON struct {
}

type XML struct {
}

func (r *JSON) recipy(dat []byte) (string, error) {
	var data recipes
	err := json.Unmarshal(dat, &data)
	if err != nil {
		log.Fatalln(err)
	}
	byte, err := xml.MarshalIndent(data, "", "    ")
	return string(byte), err
}

func (r *XML) recipy(dat []byte) (string, error) {
	var data recipes
	err := xml.Unmarshal(dat, &data)
	if err != nil {
		log.Fatalln(err)
	}
	byte, err := json.MarshalIndent(data, "", "    ")
	return string(byte), err
}

type DBReader interface {
	recipy(dat []byte) (string, error)
}

func main() {
	usage := fmt.Sprintf("Usage:\n\t./readDB -f file_name.xml\nor:\n\t./readDB -f file_name.json")
	flagFile := flag.Bool("f", false, "flag should be followed by a filename")
	flag.Parse()

	if !*flagFile || len(flag.Args()) == 0 {
		fmt.Println(usage)
		os.Exit(0)
	}

	isJson := strings.HasSuffix(flag.Args()[0], ".json")
	isXml := strings.HasSuffix(flag.Args()[0], ".xml")
	if !isJson && !isXml {
		fmt.Println(usage)
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatalln("Error: can't read from file:", err)
	}

	var db DBReader
	var js *JSON
	var xm *XML

	if isJson {
		db = js
	} else {
		db = xm
	}
	result, err := db.recipy(dat)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}