package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Contains checks if the given string is in the given array
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}

func ReadProductIDsFromFile(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var productIDs []string

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ProDuctID:") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "ProDuctID:"))
			// id, err := strconv.Atoi(idStr)
			// if err != nil {
			// 	return nil, err
			// }int

			productIDs = append(productIDs, idStr)
		}
	}

	return productIDs, nil
}
