package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/paemuri/brdoc/v2"
)

const NULL = "NULL"

func StringToBool(data string) (bool, error) {
	var isData bool

	isData, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}

	return isData, err
}

func StringToValidString(data string) string {
	if data == NULL {
		return ""
	}

	return data
}

func FormatCPFAndCNPJ(value string) (string, bool) {
	if value == NULL {
		return "", false
	}

	numericValue := regexp.MustCompile(`[^\d]`).ReplaceAllString(value, "")

	isCPF := brdoc.IsCPF(numericValue)
	isCNPJ := brdoc.IsCNPJ(numericValue)

	isValid := isCPF || isCNPJ

	return numericValue, isValid
}

func StringToDate(inputDate string) (time.Time, error) {
	if inputDate == NULL {
		return time.Time{}, nil
	}

	parsedDate, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}

func StringToFloat(data string) (*float64, error) {
	var floatData float64

	if data == NULL {
		return nil, nil
	}

	data = strings.Replace(data, ",", ".", 1)
	floatData, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	return &floatData, nil
}
