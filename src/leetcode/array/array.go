package array

import (
	"fmt"
)

type MyArrayList struct {
	data []interface{}
	size int //记录当前元素个数
}

const INIT_CAP = 1

func NewArrayList() *MyArrayList {
	return NewMyArrayListWithCap(INIT_CAP)
}

func NewMyArrayListWithCap(initCapacity int) *MyArrayList {
	return &MyArrayList{
		data: make([]interface{}, initCapacity),
		size: 0,
	}
}

//增
func (list *MyArrayList) AddLast(value interface{}) {
	cap := len(list.data)
	//查看数据容量
	if list.size == cap {
		//扩容
		list.resize(cap * 2)
	}
	//添加元素
	list.data[list.size] = value
	list.size++
}

func (list *MyArrayList) Add(index int, value interface{}) error {
	//检查是否越界
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	cap := len(list.data)
	if list.size == cap {
		list.resize(cap * 2)
	}
	//将index后面的元素向后移动 index->index+1,一直到size-1
	for i := list.size - 1; i >= index; i-- {
		list.data[i+1] = list.data[i]
	}
	list.data[index] = value
	list.size++
	return nil
}

func (list *MyArrayList) AddFirst(value interface{}) {
	list.Add(0, value)
}

func (list *MyArrayList) RemoveLast() (interface{}, error) {
	if list.size == 0 {
		return nil, fmt.Errorf("数组为空")
	}
	cap := len(list.data)
	if list.size == cap/4 {
		list.resize(cap / 2)
	}
	deleteValue := list.data[list.size-1]
	list.data[list.size-1] = nil
	list.size--
	return deleteValue, nil

}

func (list *MyArrayList) Remove(index int) (interface{}, error) {
	//检查是否越界
	if err := list.checkPositionIndex(index); err != nil {
		return nil, err
	}
	cap := len(list.data)
	if list.size == cap/4 {
		list.resize(cap / 2)
	}
	deleteValue := list.data[index]
	//将index后面的元素向前移动;index+1 -> index ；
	for i := index + 1; i < list.size; i++ {
		list.data[i-1] = list.data[i]
	}
	list.size--
	return deleteValue, nil
}

func (list *MyArrayList) RemoveFirst() (interface{}, error) {
	return list.Remove(0)
}

//查询
func (list *MyArrayList) Get(index int) (interface{}, error) {
	if err := list.checkPositionIndex(index); err != nil {
		return nil, err
	}
	return list.data[index], nil
}

//改
func (list *MyArrayList) Set(index int, value interface{}) error {
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	list.data[index] = value
	return nil
}

func (list *MyArrayList) resize(newCapacity int) {
	//重新创建1个新的内存空间，将原来的数据拷贝到新的内存空间
	newData := make([]interface{}, newCapacity)
	for i := 0; i < list.size; i++ {
		newData[i] = list.data[i]
	}
	list.data = newData

}

func (list *MyArrayList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Sprintf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *MyArrayList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}
