package validation

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

//import (
//	"github.com/go-playground/validator/v10"
//	"log"
//	"regexp"
//)
//
//func IranianMobileNumberValidation(fid validator.FieldLevel) bool {
//	value, ok := fid.Field().Interface().(string)
//	if !ok {
//		return false
//	}
//	res, err := regexp.MatchString("^[0-9]{10}$", value)
//	if err != nil {
//		log.Print(err.Error())
//	}
//	return res
//}

func IranianMobileNumberValidation(fid validator.FieldLevel) bool {
	value, ok := fid.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^\d{8,}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
