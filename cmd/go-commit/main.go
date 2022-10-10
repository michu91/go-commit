package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/michu91/go-commit/internal/services/git"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Dry run (output commit message without command execution)")
	flag.Parse()

	commit, _ := git.RunInteractiveCommit()
	commitMessage := commit.CreateCommitMessage()

	if *dryRun {
		fmt.Println(commitMessage)

		return
	}

	output, err := git.CommitCommand(commitMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
