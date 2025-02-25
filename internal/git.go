package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GitClone(currentBranch, mainBranch string) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	outStr := string(out)

	cmd = exec.Command("git", "clone", "-b", currentBranch, strings.TrimSpace(outStr), fmt.Sprintf("%s/%s", TMP_DIR, currentBranch))
	cmd.Run()
	cmd = exec.Command("git", "clone", "-b", mainBranch, strings.TrimSpace(outStr), fmt.Sprintf("%s/%s", TMP_DIR, mainBranch))
	cmd.Run()

}

func CreateWorkDir() {
	err := os.Mkdir(TMP_DIR, 0700)
	if err != nil {
		panic(err)
	}
}

func DeleteWorkDir() {
	err := os.RemoveAll(TMP_DIR)
	if err != nil {
		panic(err)
	}
}
