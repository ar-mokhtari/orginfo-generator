package preparename

import (
	"errors"
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"github.com/ar-mokhtari/orginfo-generator/pkg/convertor"
	"github.com/ar-mokhtari/orginfo-generator/pkg/validator"
	"strings"
)

func PrepareDomainName(inputDomain string) (res entity.Domain, err error) {
	if !validator.IsLuhnAlgorithm(inputDomain, env.RegexazAZ2) {
		return entity.Domain{}, errors.New(fmt.Sprintf("the value \"%s\" not support", inputDomain))
	}
	inputDomain = convertor.RemoveUnderlineFirstLast(inputDomain)
	var separator string
	if strings.Contains(inputDomain, " ") {
		separator = " "
	} else if strings.Contains(inputDomain, "_") {
		separator = "_"
	}
	switch separator {
	case " ", "_":
		var names = strings.Split(inputDomain, separator)
		var firstSlice = strings.Replace(strings.ToLower(names[0]), "_", "", -1)
		res.SnakeName, res.LowerName, res.AllLowName = firstSlice, firstSlice, firstSlice
		res.UpperName = convertor.ToUpper(firstSlice)
		for _, name := range names[1:] {
			name = strings.Replace(name, "_", "", -1)
			upper := convertor.ToUpper(name)
			res.SnakeName += "_" + convertor.ToSnake(name)
			res.LowerName += upper
			res.UpperName += upper
			res.AllLowName += strings.ToLower(name)
		}
	default:
		res.LowerName = strings.ToLower(inputDomain[:1]) + inputDomain[1:]
		res.UpperName = convertor.ToUpper(inputDomain)
		res.SnakeName = convertor.ToSnake(inputDomain)
		res.AllLowName = strings.ToLower(inputDomain)
	}
	return res, nil
}

func PrepareFieldsName(name string, fieldType string, jsonName string) (res entity.Field, err error) {
	prepareName, prepareErr := PrepareDomainName(name)
	if prepareErr != nil {
		return entity.Field{}, prepareErr
	}
	res.Name = prepareName.UpperName
	res.Type = fieldType
	res.JsonName = jsonName
	return res, nil
}
