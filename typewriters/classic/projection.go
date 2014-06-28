package classic

import (
	"regexp"
	"strings"
)

type Projection struct {
	Method string
	Type   string
	Parent *model
}

func (p Projection) MethodName() string {
	name := p.Type

	pointer := regexp.MustCompile(`^\**`)
	pointers := len(pointer.FindAllString(name, -1)[0])
	name = strings.Replace(name, "*", "", -1) + strings.Repeat("Pointer", pointers)

	slice := regexp.MustCompile(`(\[\])`)
	slices := len(slice.FindAllString(name, -1))
	name = strings.Replace(name, "[]", "", -1) + strings.Repeat("Slice", slices)

	illegal := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	name = illegal.ReplaceAllString(name, " ")

	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)

	return p.Method + strings.Title(name)
}

func (p *Projection) String() string {
	return p.MethodName()
}
