// Code generated by sqlc. DO NOT EDIT.

package datasource

import ()

type Todo struct {
	ID     int32  `json:"id"`
	Userid int32  `json:"userid"`
	Task   string `json:"task"`
	Done   bool   `json:"done"`
}

type User struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}