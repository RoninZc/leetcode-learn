package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	sum, data := make(map[uint16]struct{}), make([]int, 0)
	for _, num := range nums1 {
		sum[uint16(num)] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := sum[uint16(num)]; ok {
			data = append(data, num)
			delete(sum, uint16(num))
		}
	}
	return data
}
