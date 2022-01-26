package main

import "fmt"

// 冒泡-越到后面越大
func bubbleSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := 0; j < len(list)-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

// 冒泡2-固定最后一个元素
func bubbleSort2(list []int) []int {
	for i := len(list) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[i] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
	return list
}

// 选择-找最小的放在前面
func selectSort(sli []int) []int {
	len := len(sli)
	for i := 0; i < len-1; i++ {
		k := i // 设最小值的位置为当前
		for j := i + 1; j < len; j++ {
			if sli[k] > sli[j] {
				k = j // 记录最小值的位置
			}
		}
		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}
	}
	return sli
}

// 快排-分治
func quickSort(sli []int) []int {
	//先判断是否需要继续进行
	len := len(sli)
	if len <= 1 {
		return sli
	}
	base_num := sli[0]   //选择第一个元素作为基准
	left_sli := []int{}  //小于基准的
	right_sli := []int{} //大于基准的
	for i := 1; i < len; i++ {
		if base_num > sli[i] {
			left_sli = append(left_sli, sli[i]) // 放入左边切片
		} else {
			right_sli = append(right_sli, sli[i]) // 放入右边切片
		}
	} // 再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	left_sli = quickSort(left_sli)
	right_sli = quickSort(right_sli)
	left_sli = append(left_sli, base_num) // 合并
	return append(left_sli, right_sli...)
}

// 插入排序
func insertSort(sli []int) []int {
	len := len(sli)
	for i := 1; i < len; i++ {
		// 每次保证前面的数字有序
		tmp := sli[i] // 每次往前要插的数
		for j := i - 1; j >= 0; j-- {
			if tmp < sli[j] { // 当前数字比前面的小,则交换
				sli[j+1], sli[j] = sli[j], tmp
			} else { // 不比前面的小则退出
				break
			}
		}
	}
	return sli
}

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	var left = mergeSort(arr[0 : len(arr)/2]) // 不断地分左右
	var right = mergeSort(arr[len(arr)/2:])
	return merge(left, right) // 合并
}
func merge(left []int, right []int) []int {
	var result = make([]int, len(left)+len(right))
	var leftIndex = 0   // 左数组索引
	var rightIndex = 0  // 右数组索引
	var resultIndex = 0 // 结果索引
	// 按照从小到大的顺序往result写数据
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else if right[rightIndex] < left[leftIndex] {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}
	for leftIndex < len(left) { // 左边还有数字，加在后面
		result[resultIndex] = left[leftIndex]
		leftIndex++
		resultIndex++
	}
	for rightIndex < len(right) { // 右边还有数字，加在后面
		result[resultIndex] = right[rightIndex]
		rightIndex++
		resultIndex++
	}
	return result
}

// 堆排序-算了
func main() {
	// 排序
	list := []int{2, 6, 9, 7, 5, 3, 4, 8, 1}
	var aa = make([]int, len(list))
	copy(aa, list)
	bubbleSort2(aa)
	fmt.Println(list)             // list不变
	fmt.Println(aa)                // 冒泡排序
	fmt.Println(selectSort([]int{2, 6, 9, 7, 5, 3, 4, 8, 1})) // 选择排序
	fmt.Println(quickSort([]int{2, 6, 9, 7, 5, 3, 4, 8, 1}))  // 快速排序
	fmt.Println(insertSort([]int{2, 6, 9, 7, 5, 3, 4, 8, 1})) // 插入排序
	fmt.Println(mergeSort([]int{1, 39, 2, 9, 7, 54, 11, 8}))  // 归并排序
}
