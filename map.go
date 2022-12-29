package data

type Array []int

func (a Array) Map(callback func(int) int) Array {
	if a == nil {
		return nil
	}

	result := make(Array, len(a))
	for i, v := range a {
		result[i] = callback(v)
	}
	return result
}
