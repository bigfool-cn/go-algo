package _5_array

import (
	"errors"
	"fmt"
)

/*
Array
数组是一种线性数据结构，用连续的存储空间存储相同类型数据
数组支持随机(任意)访问
随机(任意)访问时间复杂度为O(1)
删除和插入时间复杂度：最好是0(1), 最坏是O(n), 平均是O(n)
实现插入时间复杂度O(1)的方法(前提数组不要求有序)：先将第K位的数据挪到数组末尾，再把新数据插入到第K位。
提高删除效率的方法是多次删除集中一起：记录下已经被删除的数据，每次删除操作不挪移数据，只是记录数据已被删除，
当数组没有更多的存储空间时，再统一把记录下已经被删除的数据统一从数组存储空间删除，也就是真正的删除操作。即标记清除垃圾回收算法
数组的内存寻址公式：
	一维数组：a[i]_address=base_address+i*type_size
	二维数组：二维数组假设是m*n， a[i][j]_address=base_address + (i*n+j)*type_size
	三维数组：三维数组假设是m*n*q， a[i][j][k]_address=base_address + (i*n*q + j*q + k)*type_size
*/
type Array struct {
	data   []int
	length uint
}

// NewArray 为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Len 获取数组长度
func (this *Array) Len() uint {
	return this.length
}

// 判断给定下标数组是否越界
func (this *Array) isIndexOutRange(idx uint) bool {
	if idx >= uint(cap(this.data)) {
		return true
	}
	return false
}

// Find 查找给定下标数组数据
func (this *Array) Find(idx uint) (int, error) {
	if this.isIndexOutRange(idx) {
		return 0, errors.New("out of index range")
	}
	return this.data[idx], nil
}

// Insert 插入给定下标数组数据
func (this *Array) Insert(idx uint, val int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}

	if idx != this.length && this.isIndexOutRange(idx) {
		return errors.New("out of index range")
	}

	for i := this.length; i > idx; i-- {
		this.data[i] = this.data[i-1]
	}

	this.data[idx] = val
	this.length++
	return nil
}

// InsertToTail 往数组末尾插入数据
func (this *Array) InsertToTail(val int) error {
	return this.Insert(this.Len(), val)
}

// Delete 删除给定下标数组数据
func (this *Array) Delete(idx uint) (int, error) {
	if this.isIndexOutRange(idx) {
		return 0, errors.New("out of index range")
	}

	val := this.data[idx]

	for i := idx; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}

	this.length--
	return val, nil
}

// Print 打印数组
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}

	fmt.Println(format)
}
