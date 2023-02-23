package model

import "time"

type Post struct {
	Id        string    `json:"id,omitempty"`
	Title     string    `json:"tile,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    string    `json:"author,omitempty"`
	Likes     int       `json:"likes,omitempty"`
	Dislikes  int       `json:"dislikes,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
	UpdatedAt time.Time `json:"updatedat,omitempty"`
}

func (p *Post) SetId(id string) {
	p.Id = id
}

func (p *Post) SetTitle(title string) {
	p.Title = title
}

func (p *Post) SetContent(content string) {
	p.Content = content
}

func (p *Post) SetAuthor(author string) {
	p.Author = author
}

func (p *Post) SetLike() {
	p.Likes = p.Likes + 1
}

func (p *Post) SetDislike() {
	p.Dislikes = p.Dislikes - 1
}

func (p *Post) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Post) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}
