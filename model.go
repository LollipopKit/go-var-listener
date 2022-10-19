package govarlistener

import (
	"sync"
)

type CallbackType int

const (
	OnChange = 1 << iota
	OnGet
	OnListen
	OnUnlisten
	OnError
	OnAll = 1 << iota - 1
)

type Callback[T any] struct {
	// 私有，设置后不允许更改
	fn   func(T)
	name string
	typ  CallbackType
}

type Var[T any] struct {
	// 直接以访问对象成员的方式获取值，不会触发回调函数
	Value     T
	callbacks []Callback[T]
	lock      *sync.RWMutex
}
