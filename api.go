package govarlistener

import "sync"

func NewVar[T any](value T) *Var[T] {
	return &Var[T]{
		Value:     value,
		callbacks: []Callback{},
		lock:      &sync.RWMutex{},
	}
}

func (v *Var[T]) doCallback(typ callbackType) {
	// 赋值，防止并发修改回调函数
	cbs := v.callbacks
	for idx := range cbs {
		if cbs[idx].Type & typ != 0 {
			go cbs[idx].Fn()
		}
	}
}

func (v *Var[T]) Set(value T) {
	v.Value = value
	v.doCallback(OnChange)
}

func (v *Var[T]) Get() T {
	v.doCallback(OnGet)
	return v.Value
}

// 这里不返回idx，因为可能在使用完该函数后，其他进程修改了callback list，导致idx不准确
func (v *Var[T]) IsListening(name string) bool {
	v.lock.RLock()
	defer v.lock.RUnlock()
	for idx := range v.callbacks {
		if v.callbacks[idx].Name == name {
			return true
		}
	}
	return false
}

func (v *Var[T]) Listen(callback Callback) error {
	if v.IsListening(callback.Name) {
		v.doCallback(OnError)
		return ErrSameCallbackName
	}

	v.doCallback(OnListen)

	v.lock.Lock()
	v.callbacks = append(v.callbacks, callback)
	v.lock.Unlock()
	return nil
}

func (v *Var[T]) Unlisten(name string) error {
	if !v.IsListening(name) {
		v.doCallback(OnError)
		return ErrThisNoListenName
	}

	v.doCallback(OnUnlisten)

	v.lock.Lock()
	for i := range v.callbacks {
		if v.callbacks[i].Name == name {
			v.callbacks = append(v.callbacks[:i], v.callbacks[i+1:]...)
			break
		}
	}
	v.lock.Unlock()
	return nil
}
