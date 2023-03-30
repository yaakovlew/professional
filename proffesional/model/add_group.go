package model

type AddGroup struct {
	Name        string `json:"name" binging:"required"`
	Description string `json:"description"`
}

type AllGroup struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binging:"required"`
	Description string `json:"description"`
}
