package models

type CtxString string

type Art struct {
	Name string
	Text string
}

type PageArt struct {
	Content    string
	AlterLabel string
	ArtList    []Art
}
