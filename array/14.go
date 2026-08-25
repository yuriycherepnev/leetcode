/*
Есть массив целых чисел, нужно написать функцию которая вернет массив такого же размера,
где на каждой i-ой позиции будет произведение всех элементов кроме i-го
[1, 2, 3] => [23, 13, 1*2] => [6, 3, 2]
*/
package main

func main() {

}

func MultiOther(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n)
	left := 1
	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = right
		right = nums[i]
	}
	return res
}
