package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	fmt.Println(removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
}

func removeSubfolders(folder []string) (res []string) {
	// 排序，保证顺序
	sort.Strings(folder)
	tree := make(map[string]struct{})

	for _, f := range folder {
		fatherFolder := f
		for {
			// 找到父文件夹和当前文件夹
			fatherFolder = getFatherFolder(fatherFolder)
			// fmt.Println(fatherFolder)
			// 寻找当前父文件夹是否在节点内如果存在即结束
			if _, ok := tree[fatherFolder]; ok {
				break
			}
			if fatherFolder == "" {
				tree[f] = struct{}{}
				break
			}
		}
	}
	for v := range tree {
		res = append(res, v)
	}

	return res
}

func getFatherFolder(folder string) string {
	i := strings.LastIndex(folder, "/")
	return folder[:i]
}
