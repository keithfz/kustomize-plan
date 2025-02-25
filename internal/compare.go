package internal

import (
	"fmt"
	"strings"

	godiff "codeberg.org/h7c/go-diff"
)

func Compare(newManifests, prevManifests map[string]string) {
	toCreate, toDelete := compareKeys(newManifests, prevManifests)

	createCount := len(toCreate)
	deleteCount := len(toDelete)
	updateCount := 0

	fmt.Println("Resources to create:")
	for _, v := range toCreate {
		prettyManifest := appendCharToLines(newManifests[v], "\t+ ")
		fmt.Println(prettyManifest)
		fmt.Println("")
	}
	printDivider()

	fmt.Println("Resources to delete:")
	for _, v := range toDelete {
		prettyManifest := appendCharToLines(prevManifests[v], "\t- ")
		fmt.Println(prettyManifest)
		fmt.Println("")
	}
	printDivider()

	fmt.Println("Resources to modify:")
	for k := range newManifests {

		f1 := godiff.NewFileFromBytes([]byte(newManifests[k]))
		f2 := godiff.NewFileFromBytes([]byte(prevManifests[k]))
		if f1.IsDifferentFrom(f2) {
			resource := strings.Split(k, " ")
			fmt.Printf("--- Modified %s/%s ---\n", resource[1], resource[0])
			godiff.ShowDiff(f1, f2, true)
			updateCount += 1
		}
	}
	printDivider()

	fmt.Printf("%d manifests to create, %d manifests to delete, %d manifests to update.\n", createCount, deleteCount, updateCount)

}

func compareKeys(newManifests, prevManifests map[string]string) ([]string, []string) {
	var toCreate []string
	var toDelete []string

	for k := range newManifests {
		if _, ok := prevManifests[k]; !ok {
			toCreate = append(toCreate, k)
		}
	}
	for k := range prevManifests {
		if _, ok := newManifests[k]; !ok {
			toDelete = append(toDelete, k)
		}
	}

	return toCreate, toDelete
}

func appendCharToLines(s string, char string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = char + lines[i]
	}
	return strings.Join(lines, "\n")
}

func printDivider() {
	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println()
}
