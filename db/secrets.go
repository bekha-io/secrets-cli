package db

import (
	"errors"
	"regexp"
	"time"
	"unicode/utf8"
)

type Group struct {
	ID uint
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	Secrets []Secret
}

func (g *Group) Validate() error {
	if utf8.RuneCountInString(g.Name) > 16 {
		return errors.New("Group name must be less than 16 characters")
	}

	formattedGroupName := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(g.Name, "")
	g.Name = formattedGroupName
	return nil
}

type Secret struct {
	Name string `gorm:"index:unique_secret_group,unique"`
	GroupID uint `gorm:"index:unique_secret_group,unique"`
	Group Group
	Value string
	IsHidden bool
}