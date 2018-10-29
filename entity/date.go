/* define date */
package entity

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func (date *Date) InitDate(y, m, d, h, mi, s int) {
	date.Year = y
	date.Month = m
	date.Day = d
	date.Hour = h
	date.Minute = mi
	date.Second = s
}
