package main

import (
	"fmt"
	"os"

	"github.com/keithfz/kustomize-plan/internal"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: kustomize-plan <featureBranch> <trunk> <pathToKustomization>")
		os.Exit(1)
	}

	internal.CreateWorkDir()
	defer internal.DeleteWorkDir()

	newBranch := os.Args[1]
	prevBranch := os.Args[2]
	path := os.Args[3]

	internal.GitClone(newBranch, prevBranch)

	err := internal.Build(fmt.Sprintf("%s/%s/%s", internal.TMP_DIR, prevBranch, path))
	if err != nil {
		panic(err)
	}

	prevManifests, err := internal.ParseFile(fmt.Sprintf("%s/kout.yaml", internal.TMP_DIR))
	if err != nil {
		panic(err)
	}

	err = internal.Build(fmt.Sprintf("%s/%s/%s", internal.TMP_DIR, newBranch, path))
	if err != nil {
		panic(err)
	}

	newManifests, err := internal.ParseFile(fmt.Sprintf("%s/kout.yaml", internal.TMP_DIR))
	if err != nil {
		panic(err)
	}

	internal.Compare(newManifests, prevManifests)

}
