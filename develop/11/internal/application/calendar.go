package application

import (
	"11/internal/model"
	"11/internal/storage"
	"context"
)

type Calendar struct {
	db storage.CalendarStorage
}

func NewCalendar(db storage.CalendarStorage) *Calendar {
	return &Calendar{db: db}
}

func (c *Calendar) CreateEvent(ctx context.Context, event model.Event) error {
	return c.db.Save(ctx, event)
}
func (c *Calendar) UpdateEvent(ctx context.Context, event model.Event) error {
	return c.db.Update(ctx, event)
}
func (c *Calendar) DeleteEvent(ctx context.Context, id int) error {
	return c.db.Delete(ctx, id)
}
func (c *Calendar) GetEvent(ctx context.Context, period string) ([]model.Event, error) {
	return c.db.Load(ctx, period)
}
