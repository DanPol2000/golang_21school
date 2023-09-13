package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/r3labs/diff/v3"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Itemname  string `xml:"itemname" json:"ingredient_name"`
	Itemcount string `xml:"itemcount" json:"ingredient_count"`
	Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	Name       string `xml:"name" json:"name"`
	Stovetime  string `xml:"stovetime" json:"time"`
	Ingredient []Item `xml:"ingredients>item" json:"ingredients"`
}

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Cake   `xml:"cake" json:"cake"`
}

type XML Recipe
type JSON Recipe

type DBReader interface {
	Read(content []byte) Recipe
}

func (f *XML) Read(content []byte) Recipe {
	err := xml.Unmarshal(content, f)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return Recipe(*f)
}

func (f *JSON) Read(content []byte) Recipe {
	err := json.Unmarshal(content, f)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return Recipe(*f)
}

func parseFile(reader DBReader, filename string) Recipe {
	var recipe Recipe
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	recipe = reader.Read(filecontent)
	return recipe
}

func checkFormat(filename string) string {
	if strings.HasSuffix(filename, ".xml") {
		return "xml"
	} else if strings.HasSuffix(filename, ".json") {
		return "json"
	} else {
		return ""
	}
}

func getRecipe(filename string) Recipe {
	var recipe Recipe
	format := checkFormat(filename)
	switch format {
	case "xml":
		myStruct := new(XML)
		recipe = parseFile(myStruct, filename)
	case "json":
		myStruct := new(JSON)
		recipe = parseFile(myStruct, filename)
	default:
		fmt.Fprintf(os.Stderr, "Error: invalid file extension: %s\n", filename)
		os.Exit(1)
	}
	return recipe
}

func getIngredient(path []string, recipe Recipe) string {
	cake, err := strconv.Atoi(path[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	ingredient, err := strconv.Atoi(path[3])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return recipe.Recipes[cake].Ingredient[ingredient].Itemname
}

func getCake(path []string, recipe Recipe) string {
	cake, err := strconv.Atoi(path[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return recipe.Recipes[cake].Name
}

func checkAdded(change diff.Change, recipe Recipe) {
	path := change.Path
	last := path[len(path)-1]
	switch last {
	case "Name":
		fmt.Printf("ADDED cake \"%s\"\n", change.To)
	case "Itemname":
		fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n",
			change.To, getCake(path, recipe))
	case "Itemunit":
		fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			change.To, getIngredient(path, recipe), getCake(path, recipe))
	}
}

func checkChanged(change diff.Change, recipe Recipe) {
	path := change.Path
	last := path[len(path)-1]
	switch last {
	case "Stovetime":
		fmt.Printf("CHANGED cooking time for cake \"%s\" – from \"%s\" to \"%s\"\n",
			getCake(path, recipe), change.From, change.To)
	case "Itemcount":
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" – from \"%s\" to \"%s\"\n",
			getIngredient(path, recipe), getCake(path, recipe), change.From, change.To)
	case "Itemunit":
		if change.To != "" {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" – from \"%s\" to \"%s\"\n",
				getIngredient(path, recipe), getCake(path, recipe), change.From, change.To)
		} else {
			fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n",
				getIngredient(path, recipe), getCake(path, recipe))
		}
	}
}

func checkRemoved(change diff.Change, recipe Recipe) {
	path := change.Path
	last := path[len(path)-1]
	switch last {
	case "Name":
		fmt.Printf("REMOVED cake \"%s\"\n", change.From)
	case "Itemname":
		fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n",
			getCake(path, recipe), change.From)
	case "Itemunit":
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			getIngredient(path, recipe), getCake(path, recipe), change.From)
	}
}

func compareRecipes(old Recipe, new Recipe) {
	changelog, _ := diff.Diff(old, new)
	for _, change := range changelog {
		switch change.Type {
		case diff.CREATE:
			checkAdded(change, new)
		case diff.UPDATE:
			checkChanged(change, new)
		case diff.DELETE:
			checkRemoved(change, old)
		}
	}
}

func main() {
	var oldRecipe Recipe
	var newRecipe Recipe
	oldFile := flag.String("old", "", "use old file")
	newFile := flag.String("new", "", "use new file")
	flag.Parse()

	if *oldFile != "" && *newFile != "" {
		oldRecipe = getRecipe(*oldFile)
		newRecipe = getRecipe(*newFile)
		compareRecipes(oldRecipe, newRecipe)
	} else {
		fmt.Println("Use '--old' and '--new' flags to pass arguments")
	}
}