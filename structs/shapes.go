package structs

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}
