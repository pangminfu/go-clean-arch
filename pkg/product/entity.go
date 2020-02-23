package product

import (
	"errors"
	"regexp"
)

type Product struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (p Product) IsCodeValid() error {
	re := regexp.MustCompile(`^[A-Z]{3}[0-9]{4}$`)
	isValid := re.MatchString(p.Code)
	if !isValid {
		return errors.New("invalid product code")
	}

	return nil
}
