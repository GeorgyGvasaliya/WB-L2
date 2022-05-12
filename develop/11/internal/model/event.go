package model

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        JSONTime `json:"date"`
}

//JSONTime дополнительный тип на основе time.Time
//для переопределения форматирования при маршалинге в JSON
type JSONTime time.Time

const (
	timeFormat = "02-01-2006"
)

//MarshalJSON переопределяет маршалинг для типа JSONTime
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), `"`)
	newTime, err := time.ParseInLocation(timeFormat, string(s), time.Local)
	*t = JSONTime(newTime)
	return
}
func (t JSONTime) String() string {
	return time.Time(t).Format(timeFormat)
}
