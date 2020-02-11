package inta

//go:generate mockery -name=A -case=underscore -outpkg mocks -output ./mocks

type A interface {
	GetB() B
}

type B interface {
}
