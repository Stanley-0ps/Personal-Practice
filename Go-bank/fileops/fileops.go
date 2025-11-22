package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(FileName string) (float64, error) {
	data, err := os.ReadFile(FileName)

	if err != nil {
		return 2000, errors.New("failed to find balance file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 2000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, FileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(FileName, []byte(valueText), 0644)
}
