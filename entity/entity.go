package entity

import uuid "github.com/satori/go.uuid"

type Account struct {
	ID   int       `json:"id" example:"1" format:"int64"`
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type EmptyModel struct {
	ID       int    `json:"id" example:"1" format:"int64"`
	Property string `json:"property" example:"some value"`
}

type Word struct {
	Words     string
	Meaning   string
	Sentences string
	whenToUse string
	rating    int
}
