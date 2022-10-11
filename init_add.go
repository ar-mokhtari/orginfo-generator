package generator

import (
	"fmt"
	"github.com/ar-mokhtari/orginfo/config/cli/generator/add_domain"
	"github.com/ar-mokhtari/orginfo/config/cli/generator/delete_domain"
	"github.com/ar-mokhtari/orginfo/config/cli/generator/entity"
	"github.com/ar-mokhtari/orginfo/config/env"
)

func addInit(input entity.Domain) {
	//📛delete entry domain(s) if exist
	deleteDomainErr := deletedomain.DeleteDomain(input)
	if deleteDomainErr != nil {
		fmt.Println(deleteDomainErr)
	} else {
		fmt.Println("Domain(s) deleted except `user` and `temp` ")
	}
	//dto
	dtoErr := adddomain.DomainDTOGenerator(input)
	if dtoErr != nil {
		fmt.Println(dtoErr)
	} else {
		fmt.Printf("[%s] Domain for %s created successfully \n", input.SnakeName, env.DTO)
	}
	//adapter
	adapterStorageErr := adddomain.DomainStorageGenerator(input)
	if adapterStorageErr != nil {
		fmt.Println(adapterStorageErr)
	} else {
		fmt.Printf("[%s] Domain for %s created successfully \n", input.SnakeName, env.Adapter+" -storage")
	}

	//delivery
	deliveryErr := adddomain.DomainDeliveryGenerator(input)
	if deliveryErr != nil {
		fmt.Println(deliveryErr)
	} else {
		fmt.Printf("[%s] Domain for %s created successfully \n", input.SnakeName, env.Delivery)
	}
	//entity
	entityErr := adddomain.DomainEntityGenerator(input)
	if entityErr != nil {
		fmt.Println(dtoErr)
	} else {
		fmt.Printf("[%s] Domain for %s created successfully \n", input.SnakeName, env.Entity)
	}
	//usecase
	usecaseErr := adddomain.DomainUsecaseGenerator(input)
	if usecaseErr != nil {
		fmt.Println(usecaseErr)
	} else {
		fmt.Printf("[%s] Domain for %s created successfully \n", input.SnakeName, env.Usecase)
	}
}
