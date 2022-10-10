package parser

import (
	"github.com/oriser/regroup"
)

func ParseTaskIdFromBranch(branch string) (string, error) {
	re := regroup.MustCompile(`^(feature|bugfix|hotfix)\/(?P<TaskID>[A-Z]{2,}-[0-9]+)`)

	matches, err := re.Groups(branch)

	if err != nil {
		return "", err
	}

	return matches["TaskID"], nil
}
