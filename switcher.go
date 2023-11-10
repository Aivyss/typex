package typex

type Switcher[T comparable] interface {
	Set(key T, consumer func(v *T)) error
	SetDefault(consumer func(v *T)) error
	Switch(key T) error
}
