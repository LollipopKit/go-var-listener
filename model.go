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

type Callbacks[T any] struct {
	callback []Callback[T]
	names    []string
	lock     *sync.RWMutex
}

type Var[T any] struct {
	value     T
	callbacks Callbacks[T]
}
