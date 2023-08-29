package parseJson

import (
	"encoding/json"
	"os"
)

func readFile(fileName string) (b []byte, err error) {
	b, err = os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func ParseTable(fileName string, i interface{}) error {
	b, e := readFile(fileName)
	if e != nil {
		return e
	}
	e = json.Unmarshal(b, i)
	if e != nil {
		return e
	}

	return nil
}
