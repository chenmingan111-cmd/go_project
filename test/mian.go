package main

import (
	"fmt"
)

type MySturct struct {
	name string
	age  int
}

func (m MySturct) String() string {
	return fmt.Sprintf("name is %s,age is %d", m.name, m.age)
}

func main() {
	m := MySturct{"sunset", 18}
	fmt.Println(m)
}
