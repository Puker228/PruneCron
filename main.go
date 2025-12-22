package main

import (
	"context"
	"fmt"

	"github.com/moby/moby/client"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		ctx := context.Background()
		apiClient, err := client.New(client.FromEnv)
		if err != nil {
			panic(err)
		}
		defer apiClient.Close()

		_, errorCont := apiClient.ContainerPrune(ctx, client.ContainerPruneOptions{})
		if errorCont != nil {
			fmt.Printf("Error while prune containers: %v", errorCont)
		}

		_, errorBuild := apiClient.BuildCachePrune(ctx, client.BuildCachePruneOptions{})
		if errorBuild != nil {
			fmt.Printf("Error while prune containers: %v", errorBuild)
		}
	})
	c.Start()
}
