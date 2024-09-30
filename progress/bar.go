package progress

import (
	"fmt"
)

type Progress struct {
	chars      []rune
	len        int
	percentage int
}

func (p *Progress) Init(len int) {
	p.len = len
	p.chars = make([]rune, len)
	for i := range len {
		p.chars[i] = '.'
	}
	p.percentage = 0

	fmt.Println("Reading filesystem")
}

func (p *Progress) Update(percentage int) {
	p.percentage += percentage
	for i := range p.percentage {
		p.chars[i] = '='
	}
}

func (p *Progress) Display() {
	//fmt.Print("\033[F\033[K")
	fmt.Print("\033[F")
	fmt.Print("\033[K")

	fmt.Printf("(%d/%d)", p.percentage, p.len)
	fmt.Print("[")
	for _, v := range p.chars {
		fmt.Print(string(v))
	}
	fmt.Print("]\n")
}
