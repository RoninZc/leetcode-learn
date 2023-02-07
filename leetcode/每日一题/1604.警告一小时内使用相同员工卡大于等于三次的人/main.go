package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortStrings([]string{"da", "ea", "la", "ma", "na", "xa", "ab"}))
	// fmt.Println(alertNames([]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}, []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}))
	// fmt.Println(alertNames([]string{"a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "c", "c", "c", "c", "c", "c", "c", "c", "c"}, []string{"00:37", "11:24", "14:35", "21:25", "15:48", "20:28", "07:30", "09:26", "10:32", "20:10", "19:26", "08:13", "01:08", "15:49", "02:34", "06:48", "04:33", "07:18", "00:05", "06:44", "13:33", "04:12", "03:54"}))
}

func alertNames(keyName []string, keyTime []string) (names []string) {
	res := make(map[string]([]int))

	for i := 0; i < len(keyName); i++ {
		hour := int(keyTime[i][0]-'0')*10 + int(keyTime[i][1]-'0')
		minute := int(keyTime[i][3]-'0')*10 + int(keyTime[i][4]-'0')
		res[keyName[i]] = append(res[keyName[i]], hour*60+minute)
	}

	for name, times := range res {
		sortInt(times)
		for i, time := range times[2:] {
			if time-times[i] <= 60 {
				names = append(names, name)
				break
			}
		}
	}

	sortStrings(names)
	return names
}

func sortStrings(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			point := 0
			for point < len(strs[i]) || point < len(strs[j]) {
				var iNum byte
				var jNum byte
				if point < len(strs[i]) {
					iNum = strs[i][point]
				}
				if point < len(strs[j]) {
					jNum = strs[j][point]
				}
				if iNum == jNum {
					point++
					continue
				}
				if iNum > jNum {
					strs[i], strs[j] = strs[j], strs[i]
				}
				break
			}
		}
	}
	return strs
}

func sortInt(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
