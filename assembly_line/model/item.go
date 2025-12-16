package model

import (
	"time"
)

// Item 物品介面
type Item interface {
	Process()
	GetType() string
	GetID() int
}

type Item1 struct {
	ID int
}

func (i *Item1) Process() {
	time.Sleep(100 * time.Millisecond)
}

func (i *Item1) GetType() string {
	return "Item1"
}

func (i *Item1) GetID() int {
	return i.ID
}

type Item2 struct {
	ID int
}

func (i *Item2) Process() {
	time.Sleep(200 * time.Millisecond)
}

func (i *Item2) GetType() string {
	return "Item2"
}

func (i *Item2) GetID() int {
	return i.ID
}

type Item3 struct {
	ID int
}

func (i *Item3) Process() {
	time.Sleep(300 * time.Millisecond)
}

func (i *Item3) GetType() string {
	return "Item3"
}

func (i *Item3) GetID() int {
	return i.ID
}

func NewItem1(id int) *Item1 {
	return &Item1{ID: id}
}

func NewItem2(id int) *Item2 {
	return &Item2{ID: id}
}

func NewItem3(id int) *Item3 {
	return &Item3{ID: id}
}
