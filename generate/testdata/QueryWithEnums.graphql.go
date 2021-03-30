package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
)

type QueryWithEnumsResponse struct {
	User QueryWithEnumsUser `json:"user"`
}

type QueryWithEnumsUser struct {
	Roles []QueryWithEnumsUserRolesRole `json:"roles"`
}

type QueryWithEnumsUserRolesRole string

const (
	QueryWithEnumsUserRolesRoleStudent QueryWithEnumsUserRolesRole = "STUDENT"
	QueryWithEnumsUserRolesRoleTeacher QueryWithEnumsUserRolesRole = "TEACHER"
)

func QueryWithEnums(
	client graphql.Client,
) (*QueryWithEnumsResponse, error) {
	var retval QueryWithEnumsResponse
	err := client.MakeRequest(
		nil,
		"QueryWithEnums",
		`
query QueryWithEnums {
	user {
		roles
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}