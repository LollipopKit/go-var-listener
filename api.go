package govarlistener

import "sync"

func NewVar[T any](value T) *Var[T] {
	return &Var[T]{
		value: value,
		callbacks: Callbacks[T]{
			callback: []Callback[T]{},
			names:    []string{},
			lock:     &sync.RWMutex{},
		},
	}
}

func (v *Var[T]) Set(value T) {
	v.value = value

	// 赋值，防止并发修改回调函数
	cbs := v.callbacks.callback
	for _, callback := range cbs {
		if callback.typ == OnChange || callback.typ == OnBoth {
			go callback.fn(value)
		}
	}
}

func (v *Var[T]) Get() T {
	value := v.value
	cbs := v.callbacks.callback
	for _, cb := range cbs {
		if cb.typ == OnGet || cb.typ == OnBoth {
			go cb.fn(value)
		}
	}

	return value
}

func (v *Var[T]) HaveListen(name string) bool {
	v.callbacks.lock.RLock()
	defer v.callbacks.lock.RUnlock()
	return Contains(v.callbacks.names, name)
}

func (v *Var[T]) Listen(callback Callback[T]) error {
	if v.HaveListen(callback.name) {
		return ErrSameCallbackName
	}

	v.callbacks.lock.Lock()
	v.callbacks.callback = append(v.callbacks.callback, callback)
	v.callbacks.names = append(v.callbacks.names, callback.name)
	v.callbacks.lock.Unlock()
	return nil
}

func (v *Var[T]) Unlisten(name string) error {
	if !v.HaveListen(name) {
		return ErrThisNoListenName
	}

	v.callbacks.lock.Lock()
	for i, cb := range v.callbacks.callback {
		if cb.name == name {
			v.callbacks.callback = append(v.callbacks.callback[:i], v.callbacks.callback[i+1:]...)
			break
		}
	}
	for i, cbName := range v.callbacks.names {
		if cbName == name {
			v.callbacks.names = append(v.callbacks.names[:i], v.callbacks.names[i+1:]...)
			break
		}
	}
	v.callbacks.lock.Unlock()
	return nil
}
