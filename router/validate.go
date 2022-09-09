package rout

import (
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
)

func ValidateData(userURL string, newURL string) (string, error) {

	val1 := govalidator.IsAlphanumeric(newURL)
	if val1 == false {
		log.Println("newurl is not url")
		return "", fmt.Errorf("newurl is not url")
	}

	val2 := govalidator.IsURL(userURL)
	if val2 == false {
		log.Println("userurl is not url")
		return "", fmt.Errorf("userurl is not url")
	}

	if newURL == "" || userURL == "" {
		log.Println("newurl or userurl is not url")
		return "", fmt.Errorf("newurl or userurl is not url")
	}

	if strings.HasPrefix(userURL, "www.") {
		newUserURL := strings.ReplaceAll(userURL, "www.", "https://")
		return newUserURL, nil

	}
	if strings.HasPrefix(userURL, "") {
		newUserURL := strings.Replace(userURL, "", "https://", +1)
		return newUserURL, nil
	}

	return "", nil
}
