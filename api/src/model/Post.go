package model

import "errors"

type Post struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"tile,omitempty"`
	Content   string `json:"content,omitempty"`
	Author    string `json:"author,omitempty"`
	Likes     int    `json:"likes,omitempty"`
	Dislikes  int    `json:"dislikes,omitempty"`
	CreatedAt string `json:"createdat,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty"`
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

func (p *Post) SetCreatedAt(createdAt string) {
	p.CreatedAt = createdAt
}

func (p *Post) SetUpdatedAt(updatedAt string) {
	p.UpdatedAt = updatedAt
}

func ValidatePost(p Post) error {
	if len(p.Title) <= 0 {
		return errors.New("invalid 'title' field")
	} else if len(p.Content) <= 0 {
		return errors.New("invalid 'content' field")
	}

	return nil
}
