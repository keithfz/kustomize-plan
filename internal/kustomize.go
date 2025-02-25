package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Build(path string) error {
	cmd := exec.Command("kustomize", "build", "--enable-helm", "--load-restrictor", "LoadRestrictionsNone", path)
	file, err := os.Create(fmt.Sprintf("%s/%s", TMP_DIR, "kout.yaml"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
