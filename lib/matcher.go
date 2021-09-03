package lib

import (
	"context"
	"fmt"

	"github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db"
)

func MatchDisease(disease string) string {
	message := fmt.Sprintf("test")
	fmt.Println("gaming")
	return message
}

func CreateDisease(name string, part string, symptoms []string) error { 
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	
	newDisease, err := client.Disease.CreateOne(
		db.Disease.Name.Set(name),
		db.Disease.Part.Set(part),
		db.Disease.Symptoms.Set(symptoms),
	).Exec(ctx)

	if err != nil {
		return err
	}

	fmt.Println(newDisease)
	return nil
}

