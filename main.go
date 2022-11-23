package main

import (
	"fmt"

	"github.com/lukasjarosch/skipper"
	"github.com/spf13/afero"
)

func main() {

	fileSystem := afero.NewOsFs()
	targetName := "dev"

	inventory, err := skipper.NewInventory(fileSystem, "inventory/classes", "inventory/targets", "inventory/secrets")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := inventory.Data(targetName, nil, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	templater, err := skipper.NewTemplater(fileSystem, "templates", "deploy/"+targetName, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	skipperConfig, err := inventory.GetSkipperConfig(targetName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = templater.ExecuteAll(skipper.DefaultTemplateContext(data, targetName), false, skipperConfig.Renames)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
