package variables

import (
	"os"
	"strconv"
)

func Get(name string) (value string, err error) {
	value = os.Getenv(name)
	return
}

func GetInt(name string) (value int, err error) {
	valueString, err := Get(name)
	if err != nil {
		return
	}

	value, err = strconv.Atoi(valueString)
	return
}

func GetInt64(name string) (value int64, err error) {
	valueString, err := Get(name)
	if err != nil {
		return
	}

	value, err = strconv.ParseInt(valueString, 10, 64)
	return
}
