package govarlistener

import "sync"

func NewVar[T any](value T) *Var[T] {
	return &Var[T]{
		value:     value,
		callbacks: []Callback[T]{},
		lock:      &sync.RWMutex{},
	}
}

func (v *Var[T]) Set(value T) {
	v.value = value

	// 赋值，防止并发修改回调函数
	cbs := v.callbacks
	for _, callback := range cbs {
		if callback.typ == OnChange || callback.typ == OnBoth {
			go callback.fn(value)
		}
	}
}

func (v *Var[T]) Get() T {
	value := v.value
	cbs := v.callbacks
	for _, cb := range cbs {
		if cb.typ == OnGet || cb.typ == OnBoth {
			go cb.fn(value)
		}
	}

	return value
}

// 这里不返回idx，因为可能在使用完该函数后，其他进程修改了callback list，导致idx不准确
func (v *Var[T]) HaveListen(name string) bool {
	v.lock.RLock()
	defer v.lock.RUnlock()
	for _, cb := range v.callbacks {
		if cb.name == name {
			return true
		}
	}
	return false
}

func (v *Var[T]) Listen(callback Callback[T]) error {
	if v.HaveListen(callback.name) {
		return ErrSameCallbackName
	}

	v.lock.Lock()
	v.callbacks = append(v.callbacks, callback)
	v.lock.Unlock()
	return nil
}

func (v *Var[T]) Unlisten(name string) error {
	if !v.HaveListen(name) {
		return ErrThisNoListenName
	}

	v.lock.Lock()
	for i, cb := range v.callbacks {
		if cb.name == name {
			v.callbacks = append(v.callbacks[:i], v.callbacks[i+1:]...)
			break
		}
	}
	v.lock.Unlock()
	return nil
}
