package main

import (
	"context"
	"fmt"
	"test/models"
)

func main() {
	fmt.Println("Hello, world!")
	modelClient := models.NewModelClient("admin")
	if modelClient == nil {
		fmt.Println("modelClient is nil")
		return
	}
	ctx := context.Background()

	modelDataI, err := modelClient.GetById(ctx, 123)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if modelDataI != nil {
		modelClient.ReturnData(ctx, modelDataI, nil)
	}

	modelData := modelDataI.(models.Admin)
	modelData.Save(ctx)
	modelData.ReturnData(ctx, modelData, nil)
	fmt.Println("modelData:", modelData)

}
