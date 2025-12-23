package main

import (
	"context"
	"fmt"

	"github.com/moby/moby/client"
)

func main() {
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

	f := make(client.Filters).Add("dangling", "true")
	_, errorImage := apiClient.ImagePrune(ctx, client.ImagePruneOptions{
		Filters: f,
	})

	if errorImage != nil {
		fmt.Printf("Error while prune containers: %v", errorImage)
	}

	_, errorVolume := apiClient.VolumePrune(ctx, client.VolumePruneOptions{})
	if errorVolume != nil {
		fmt.Printf("Error while prune containers: %v", errorVolume)
	}
}
