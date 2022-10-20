package govarlistener

import (
	"sync"
)

type callbackType int

const (
	OnChange = 1 << iota
	OnGet
	OnListen
	OnUnlisten
	OnError
	OnAll = 1<<iota - 1
)

type Callback struct {
	Fn   func()
	Name string
	Type callbackType
}

type Var[T any] struct {
	// 直接以访问对象成员的方式获取值，不会触发回调函数
	Value     T
	callbacks []Callback
	lock      *sync.RWMutex
}
