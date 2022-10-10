package parser

import (
	"github.com/oriser/regroup"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestParseTaskIdFromBranchSuccess(t *testing.T) {
	examples := map[string]string{
		"feature/KOD-1102-some-feature":     "KOD-1102",
		"bugfix/DAOS-001-another-feature":   "DAOS-001",
		"feature/TASK-1-git-commit-wrapper": "TASK-1",
		"hotfix/FEA-1-some-hotfix-DDD-123":  "FEA-1",
	}

	for example, expected := range examples {
		result, err := ParseTaskIdFromBranch(example)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestParseTaskIdFromBranchNoFound(t *testing.T) {
	result, err := ParseTaskIdFromBranch("feature/without-task-id")

	assert.IsType(t, &regroup.NoMatchFoundError{}, err)
	assert.Equal(t, "", result)
}
