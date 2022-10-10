package git

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/michu91/go-commit/internal/model"
	"github.com/michu91/go-commit/internal/services/parser"
)

func RunInteractiveCommit() (*model.Commit, error) {
	var (
		taskID string
		commit model.Commit
	)

	branch, err := GetCurrentBranch()
	if err == nil {
		taskID, err = parser.ParseTaskIdFromBranch(branch)
	}

	var questions = []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Select type of commit message:",
				Options: []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore"},
				Default: "feat",
			},
			Validate: survey.Required,
		},
		{
			Name:   "scope",
			Prompt: &survey.Input{Message: "Scope (optional, enter to skip): "},
		},
		{
			Name:     "shortDescription",
			Prompt:   &survey.Input{Message: "Short description: "},
			Validate: survey.Required,
		},
		{
			Name:   "longDescription",
			Prompt: &survey.Input{Message: "Long description (optional, enter to skip): "},
		},
		{
			Name:   "taskId",
			Prompt: &survey.Input{Message: "Issue references: ", Default: taskID},
		},
		{
			Name: "isBreakingChange",
			Prompt: &survey.Input{
				Message: "Is breaking change <y/n>?: ",
				Default: "n",
			},
			Validate: func(val interface{}) error {
				if str := val.(string); str != "y" && str != "n" {
					return fmt.Errorf("you entered '%s', not 'y/n'", str)
				}

				return nil
			},
			Transform: func(ans interface{}) interface{} {
				val := ans.(string)

				return val == "y"
			},
		},
	}

	err = survey.Ask(questions, &commit)
	if err != nil {
		return nil, err
	}

	return &commit, nil
}
