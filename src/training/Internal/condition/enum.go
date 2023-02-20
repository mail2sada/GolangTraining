package main

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednsday = 4
	Thursady = iota
	Friday
	Saturday
)

func main() {

	var weekDay WeekDay = Sunday

	switch weekDay {

	case Sunday:
	case Monday:
	case Tuesday:
	default:
	}
}
