package govarlistener

import "errors"

var (
	ErrSameCallbackName = errors.New("callback with same name already exists")
	ErrThisNoListenName = errors.New("this callback name does not exist")
)
