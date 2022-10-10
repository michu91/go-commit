package model

import (
	"fmt"
	"strings"
)

type Commit struct {
	Type             string
	Scope            string
	IsBreakingChange bool
	ShortDescription string
	LongDescription  string
	TaskId           string
	Footer           string
}

func (c *Commit) CreateCommitMessage() string {
	var builder strings.Builder

	builder.WriteString(c.Type)

	if c.IsBreakingChange {
		builder.WriteString("!")
	}

	if c.Scope != "" {
		builder.WriteString(fmt.Sprintf("(%s)", c.Scope))
	}

	builder.WriteString(": ")
	builder.WriteString(c.ShortDescription)

	if c.LongDescription != "" {
		builder.WriteString("\n")
		builder.WriteString(c.LongDescription)
	}

	if c.TaskId != "" {
		builder.WriteString("\n")
		builder.WriteString(fmt.Sprintf("Refs: #%s", c.TaskId))
	}

	return builder.String()
}
