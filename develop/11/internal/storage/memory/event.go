package memory

import (
	"11/internal/model"
	"context"
	"time"
)

func (m *MemoryDB) Save(_ context.Context, event model.Event) error {
	m.Lock()
	defer m.Unlock()
	event.ID = m.nextID()
	m.data[event.ID] = event
	return nil
}
func (m *MemoryDB) Update(_ context.Context, event model.Event) error {
	_, ok := m.data[event.ID]
	if !ok {
		return model.ErrNotFound
	}
	m.data[event.ID] = event
	return nil
}
func (m *MemoryDB) Delete(_ context.Context, id int) error {
	_, ok := m.data[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(m.data, id)
	return nil
}
func (m *MemoryDB) Load(_ context.Context, interval string) ([]model.Event, error) {
	switch interval {
	case "day":
		nowDay := time.Now()

		todayStart := time.Date(nowDay.Year(), nowDay.Month(), nowDay.Day(), 0, 0, 0, 0, nowDay.Location())
		todayEnd := todayStart.AddDate(0, 0, 1).Add(time.Nanosecond * -1)

		return m.getBetween(todayStart, todayEnd)
	case "week":
		nowDay := time.Now()
		firstWeekDay := nowDay

		for firstWeekDay.Weekday() != time.Monday {
			firstWeekDay = firstWeekDay.AddDate(0, 0, -1)
		}
		firstWeekDay = time.Date(firstWeekDay.Year(), firstWeekDay.Month(),
			firstWeekDay.Day(), 0, 0, 0, 0, firstWeekDay.Location())
		lastWeekDay := firstWeekDay.AddDate(0, 0, 7).Add(time.Nanosecond * -1)
		return m.getBetween(firstWeekDay, lastWeekDay)
	case "month":
		nowDay := time.Now()

		monthStart := time.Date(nowDay.Year(), nowDay.Month(), 1, 0, 0, 0, 0, nowDay.Location())
		monthEnd := monthStart.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
		return m.getBetween(monthStart, monthEnd)

	default:
		return nil, model.ErrInvalidInterval
	}
}

//getBetween возвращает все события, запланированные в промежутке между временем cтарта и окончания
//Принимает: время старта и время окончания.
//Возвращает: слайс событий и ошибку получения.
//Конкурентно безопасный метод.
func (m *MemoryDB) getBetween(start time.Time, end time.Time) ([]model.Event, error) {
	result := make([]model.Event, 0)
	m.RLock()
	defer m.RUnlock()
	for _, v := range m.data {

		if inTimeSpan(start, end, time.Time(v.Date)) {
			result = append(result, v)
		}

	}
	return result, nil
}

//inTimeSpan проверяет нахождение времени в заданном промежутке.
//Принимает: время старта промежутка, время окончания промежутка и проверяемое время.
//Возвращает: булевский результат нахождения.
func inTimeSpan(start, end, check time.Time) bool {
	if start.After(end) {
		start, end = end, start
	}

	return !check.Before(start) && !check.After(end)
}
