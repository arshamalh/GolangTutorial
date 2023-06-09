package cmap

// Custom map
type CMap[T any] map[string]T

func (cmap CMap[T]) Set(key string, value T) {
	cmap[key] = value
}

func (cmap CMap[T]) Get(key string) T {
	return cmap[key]
}

func (cmap CMap[T]) Pop(key string) T {
	value := cmap.Get(key)
	delete(cmap, key)
	return value
}

func (cmap CMap[T]) Keys() []string {
	keys := make([]string, 0)
	for key := range cmap {
		keys = append(keys, key)
	}
	return keys
}

func (cmap CMap[T]) Values() []T {
	values := make([]T, 0)
	for _, value := range cmap {
		values = append(values, value)
	}
	return values
}

type Pair[Y any] struct {
	Key   string
	Value Y
}

func (cmap CMap[T]) Pairs() []Pair[T] {
	pairs := make([]Pair[T], 0)
	for key, value := range cmap {
		new_pair := Pair[T]{
			Key:   key,
			Value: value,
		}
		pairs = append(pairs, new_pair)
	}
	return pairs
}
