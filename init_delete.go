package generator

import (
	"fmt"
	"github.com/ar-mokhtari/orginfo-generator/delete_domain"
	"github.com/ar-mokhtari/orginfo-generator/entity"
)

func deleteInit(inputDomain entity.Domain) {
	err := deletedomain.DeleteDomain(inputDomain)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s Domain deleted successfully \n", inputDomain.SnakeName)
	}
}
