package generator

import (
	"fmt"
	"github.com/ar-mokhtari/orginfo/config/cli/generator/delete_domain"
	"github.com/ar-mokhtari/orginfo/config/cli/generator/entity"
)

func deleteInit(inputDomain entity.Domain) {
	err := deletedomain.DeleteDomain(inputDomain)
	if err != nil {
		fmt.Println(err)
	}
}
