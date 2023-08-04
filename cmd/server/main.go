package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdikhodaparast/vgang-challenge/cmd"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/service/utils"
)

func main() {
	ctx := context.Background()
	startTime := time.Now()
	factory, err := cmd.NewFactory()
	if err != nil {
		log.Panicln("error creating factory", err)
	}
	doneChannel := factory.DoneChannel
	for _, runner := range factory.PipelineStages {
		go func(runner model.Runner) {
			err := runner.Run(ctx)
			if err != nil {
				log.Panicln("error running runner", err)
			}
		}(runner)
	}
	<-doneChannel
	log.Println("total time:", time.Since(startTime))

	filePath := "result.txt"
	productIDs, err := utils.ReadProductIDsFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	shorter := utils.New()
	app := fiber.New()
	shorter.CreateIntMap(productIDs)
	app.Post("/createShortUrl", shorter.CreateURL)
	app.Get("/useShortUrl/:shortURL", shorter.Redirect)
	app.Get("/getAllShortUrl", shorter.GetAll)
	//Port Can get from config file
	port := 8088
	fmt.Printf("Server is listening on port %d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
