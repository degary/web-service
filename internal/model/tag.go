package model

import "github.com/degary/web-service/pkg/app"

type Tag struct {
	*Model
	Name string `json:"name"`
	State string `json:"state"`
}

func (a Tag)TableName()string {
	return "blog_tag"
}


type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}