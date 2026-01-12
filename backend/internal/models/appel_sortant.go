package models

type AppelSortant[T any] struct {
	Code    int
	Message string
	Sortie  T
}
