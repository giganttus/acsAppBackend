package models

import "time"

type Posts struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Author   string    `json:"author"`
	Category string    `json:"category,omitempty"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	Featured bool      `json:"featured"`
	ILink    string    `json:"ilink"`
	Subtitle string    `json:"subtitle"`
	Tags     string    `json:"tags,omitempty"`
	Title    string    `json:"title"`
}

type PostsPreview struct {
	ID       int       `json:"id"`
	Author   string    `json:"author"`
	Category string    `json:"category,omitempty"`
	Date     time.Time `json:"date"`
	Featured bool      `json:"featured"`
	ILink    string    `json:"ilink"`
	Subtitle string    `json:"subtitle"`
	Tags     string    `json:"tags,omitempty"`
	Title    string    `json:"title"`
}
