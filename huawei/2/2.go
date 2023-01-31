package main

import "fmt"

func main() {
	count, num := 0, 0
	_, _ = fmt.Scanln(&count)
	counter := NewCounter()
	for i := 0; i < count; i++ {
		_, _ = fmt.Scanln(&num)
		counter.Add(num)
	}
	for _, i := range counter.GetData() {
		fmt.Println(i)
	}
}

func NewCounter() Counter {
	return Counter{
		Set:  make(map[int]struct{}),
		Data: make([]int, 0),
	}
}

type Counter struct {
	Set  map[int]struct{}
	Data []int
}

func (c *Counter) Add(num int) {
	if _, ok := c.Set[num]; ok {
		return
	}
	c.Set[num] = struct{}{}
	c.Data = append(c.Data, num)
}

func (c *Counter) GetData() []int {
	for i := 0; i < len(c.Data); i++ {
		for j := i + 1; j < len(c.Data); j++ {
			if c.Data[i] > c.Data[j] {
				c.Data[i], c.Data[j] = c.Data[j], c.Data[i]
			}
		}
	}

	return c.Data
}
