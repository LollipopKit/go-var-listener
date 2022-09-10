package govarlistener

import (
	"sync"
)

type CallbackType int

const (
	OnChange CallbackType = iota
	OnGet
	OnBoth
)

type Callback[T any] struct {
	fn   func(T)
	name string
	typ  CallbackType
}

type Var[T any] struct {
	value     T
	callbacks []Callback[T]
	lock      *sync.RWMutex
}
