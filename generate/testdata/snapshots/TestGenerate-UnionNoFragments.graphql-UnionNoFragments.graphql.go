package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// UnionNoFragmentsQueryRandomLeafArticle includes the requested fields of the GraphQL type Article.
type UnionNoFragmentsQueryRandomLeafArticle struct {
	Typename string `json:"__typename"`
}

// UnionNoFragmentsQueryRandomLeafLeafContent includes the requested fields of the GraphQL interface LeafContent.
//
// UnionNoFragmentsQueryRandomLeafLeafContent is implemented by the following types:
// UnionNoFragmentsQueryRandomLeafArticle
// UnionNoFragmentsQueryRandomLeafVideo
//
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type UnionNoFragmentsQueryRandomLeafLeafContent interface {
	implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *UnionNoFragmentsQueryRandomLeafArticle) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface UnionNoFragmentsQueryRandomLeafLeafContent.
func (v *UnionNoFragmentsQueryRandomLeafArticle) GetTypename() string { return v.Typename }

func (v *UnionNoFragmentsQueryRandomLeafVideo) implementsGraphQLInterfaceUnionNoFragmentsQueryRandomLeafLeafContent() {
}

// GetTypename is a part of, and documented with, the interface UnionNoFragmentsQueryRandomLeafLeafContent.
func (v *UnionNoFragmentsQueryRandomLeafVideo) GetTypename() string { return v.Typename }

func __unmarshalUnionNoFragmentsQueryRandomLeafLeafContent(v *UnionNoFragmentsQueryRandomLeafLeafContent, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(UnionNoFragmentsQueryRandomLeafArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(UnionNoFragmentsQueryRandomLeafVideo)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for UnionNoFragmentsQueryRandomLeafLeafContent: "%v"`, tn.TypeName)
	}
}

// UnionNoFragmentsQueryRandomLeafVideo includes the requested fields of the GraphQL type Video.
type UnionNoFragmentsQueryRandomLeafVideo struct {
	Typename string `json:"__typename"`
}

// UnionNoFragmentsQueryResponse is returned by UnionNoFragmentsQuery on success.
type UnionNoFragmentsQueryResponse struct {
	RandomLeaf UnionNoFragmentsQueryRandomLeafLeafContent `json:"-"`
}

func (v *UnionNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {

	type UnionNoFragmentsQueryResponseWrapper UnionNoFragmentsQueryResponse

	var firstPass struct {
		*UnionNoFragmentsQueryResponseWrapper
		RandomLeaf json.RawMessage `json:"randomLeaf"`
	}
	firstPass.UnionNoFragmentsQueryResponseWrapper = (*UnionNoFragmentsQueryResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.RandomLeaf
		raw := firstPass.RandomLeaf
		err = __unmarshalUnionNoFragmentsQueryRandomLeafLeafContent(
			target, raw)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnionNoFragmentsQuery(
	client graphql.Client,
) (*UnionNoFragmentsQueryResponse, error) {
	var retval UnionNoFragmentsQueryResponse
	err := client.MakeRequest(
		nil,
		"UnionNoFragmentsQuery",
		`
query UnionNoFragmentsQuery {
	randomLeaf {
		__typename
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

