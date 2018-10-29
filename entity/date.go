/* define date */
package entity

import (
	"regexp"
	"strconv"
)

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

func (date *Date) InitDate(y, m, d, h, mi int) {
	date.Year = y
	date.Month = m
	date.Day = d
	date.Hour = h
	date.Minute = mi
}

func CheckDateValid(d Date) bool {
	checkM := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if d.Year < 1000 || d.Year > 9999 ||
		d.Month < 1 || d.Month > 12 ||
		d.Day < 1 || d.Day > 31 ||
		d.Hour < 0 || d.Hour >= 24 ||
		d.Minute < 0 || d.Minute >= 60 {
		return false
	}
	if d.Year%400 == 0 || (d.Year%4 == 0 && d.Year%100 != 0) {
		checkM[1] = 29
	}
	if d.Day > checkM[d.Month-1] {
		return false
	}
	return true
}

// 0000-00-00/00:00
func DateToStr(d Date) string {
	str := ExpendStr(IntToStr(d.Year), 4) + "-" + ExpendStr(IntToStr(d.Month), 2) + "-" +
		ExpendStr(IntToStr(d.Day), 2) + "/" + ExpendStr(IntToStr(d.Hour), 2) + ":" +
		ExpendStr(IntToStr(d.Minute), 2)
	return str
}

func StrToDate(s string) Date {
	match, _ := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2}/[0-9]{2}:[0-9]{2}", s)
	if match {
		d := Date{
			StrToInt(s[0:4]),
			StrToInt(s[5:7]),
			StrToInt(s[8:10]),
			StrToInt(s[11:13]),
			StrToInt(s[14:16]),
		}
		if CheckDateValid(d) {
			return d
		} else {
			return Date{0, 0, 0, 0, 0}
		}
	} else {
		return Date{0, 0, 0, 0, 0}
	}
}

func StrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}

func IntToStr(n int) string {
	str := strconv.Itoa(n)
	return str
}

func ExpendStr(s string, n int) string {
	for i := len(s); i < n; i++ {
		s = "0" + s
	}
	return s
}
