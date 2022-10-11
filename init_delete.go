package generator

import (
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/entity"
)

func deleteInit(inputDomain entity.Domain) {
	err := deletedomain.DeleteDomain(inputDomain)
	if err != nil {
		fmt.Println(err)
	}
}
