package main

import (
	"context"
	"fmt"
	"os/exec"

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

	cmd := exec.Command("docker", "image", "prune", "-a", "-f")
	imageError := cmd.Run()
	if imageError != nil {
		fmt.Println("Error while prune images")
	}

	_, errorVolume := apiClient.VolumePrune(ctx, client.VolumePruneOptions{})
	if errorVolume != nil {
		fmt.Printf("Error while prune containers: %v", errorVolume)
	}
}
