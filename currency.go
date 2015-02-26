package money

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	currencies map[string]map[string]interface{}
)

func get(attribute string, currency string) interface{} {
	return currencies[currency][attribute]
}

func init() {
	file, err := ioutil.ReadFile("./data.json")

	if err != nil {
		log.Println("[money] %s", err)
	}

	err = json.Unmarshal(file, &currencies)

	if err != nil {
		log.Println("[money] %s", err)
	}
}
