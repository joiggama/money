package money

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	currencies     map[string]map[string]interface{}
	currenciesFile = os.Getenv("GOPATH") + "/src/github.com/joiggama/money/currencies.json"
)

func init() {
	file, err := ioutil.ReadFile(currenciesFile)

	if err != nil {
		log.Fatalf("[money] %s\n", err)
	}

	err = json.Unmarshal(file, &currencies)

	if err != nil {
		log.Fatalf("[money] %s\n", err)
	}
}
