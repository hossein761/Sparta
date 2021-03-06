package iam

import (
	gocf "github.com/hossein761/go-cloudformation"
)

// PolicyStatement represents an entry in an IAM policy document
type PolicyStatement struct {
	Effect   string
	Action   []string
	Resource *gocf.StringExpr
}
