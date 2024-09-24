// models/repos.go

package models

type Repo struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Org  string `json:"org"`
	Repo string `json:"repo" gorm:"unique"`
	Name string `json:"name"`
}

type AddRepoInput struct {
	Org  string `json:"org" binding:"required"`
	Repo string `json:"repo" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateRepoInput struct {
	Org  string `json:"org"`
	Repo string `json:"repo"`
	Name string `json:"name"`
}
