package main

import "fmt"

func main() {
	ECounter := ErrCounter{
		Data:  make(map[string]int),
		Order: make([]string, 0),
	}
	for {
		file, line := "", 0
		ok, _ := fmt.Scanf("%s%d", &file, &line)
		if ok == 0 {
			break
		}
		ECounter.store(fmt.Sprintf("%s %d", file, line))
	}
	ECounter.OutPut()
}

type ErrCounter struct {
	Data  map[string]int
	Order []string
}

func (e *ErrCounter) OutPut() {
	var is bool
	for i := 0; i < len(e.Order); i++ {
		is = true
		for j := 0; j < len(e.Order)-1; j++ {
			if e.Data[e.Order[j]] < e.Data[e.Order[j+1]] {
				e.Order[j+1], e.Order[j] = e.Order[j], e.Order[j+1]
				is = false
			}
		}
		if is {
			break
		}
	}

	count := 0
	for _, file := range e.Order {
		if count == 8 {
			return
		}
		filename := file
		for i := 0; i < len(filename); i++ {
			if filename[i] == ' ' {
				if i > 16 {
					filename = filename[i-16:]
				}
			}
		}
		fmt.Printf("%s %d\n", filename, e.Data[file])
		count++
	}
}

func (e *ErrCounter) store(filename string) {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '\\' {
			filename = filename[i+1:]
		}
	}

	if _, ok := e.Data[filename]; !ok {
		e.Data[filename] = 1
		e.Order = append(e.Order, filename)
	} else {
		e.Data[filename]++
	}
}
