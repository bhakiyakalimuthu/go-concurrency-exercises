package main

type req struct {
	a *a
	b *b
}

func NewReq(a *a, b *b) *req {
	return &req{
		a: a,
		b: b,
	}
}

type a struct {
	i, y int
}
type b struct {
	i, y int
}
