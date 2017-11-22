package containers

type Array struct {
	size  int
	datas []interface{}
}

func NewArray(size int) *Array {
	instance := &Array{
		size:  size,
		datas: make([]interface{}, 0, size),
	}
	return instance
}

func (arr *Array) GetLength() int {
	return len(arr.datas)
}

func (arr *Array) Push(data interface{}) bool {
	if len(arr.datas) < arr.size {
		arr.datas = append(arr.datas, data)
		return true
	}
	return false
}

func (arr *Array) Pop() interface{} {
	if len(arr.datas) > 0 {
		data := arr.datas[len(arr.datas)-1]
		arr.datas = arr.datas[:len(arr.datas)-1]
		return data
	}
	return nil
}

func (arr *Array) More(size int) {
	arr.size += size
}
