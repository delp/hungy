package hungy

import (
	"fmt"
	"time"
)

type Stockpile struct {
	Items []Item
}

func (s *Stockpile) AddItems(i ...Item) {
	s.Items = append(s.Items, i...)
}

func (s *Stockpile) GetItems() []Item {
	return s.Items
}

type Item struct {
	Description string
	Acquired    time.Time
	Expires     time.Time
	BestBy      time.Time
	Tags        []string
}

type Taggable interface {
	Tag(tags ...string)
	GetTags() []string
}

func (i *Item) Tag(tags ...string) {
	i.Tags = append(i.Tags, tags...)
}

func (i *Item) GetTags() []string {
	return i.Tags
}

type Perishable interface {
	Expires() time.Time
	Acquired() time.Time
}

func (i *Item) ExpiresAt() time.Time {
	return i.Expires
}

func (i *Item) AcquiredAt() time.Time {
	return i.Acquired
}

