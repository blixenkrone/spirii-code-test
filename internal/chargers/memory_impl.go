package chargers

import (
	"blixenkrone/spirii/storage"
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrRecordNotFound = errors.New("foo record not found")
)

type FooDB struct {
	cache map[string]MeterReading
}

type MeterReading struct {
	Timestamp       time.Time
	MeterID         string
	ConsumerID      string
	MeterReadingVal int
}

// Read implements storage.ProjectionReadWriter.
func (f *FooDB) Read(ctx context.Context, id string) (MeterReading, error) {
	if v, ok := f.cache[id]; ok {
		return v, nil
	}
	return MeterReading{}, fmt.Errorf("error finding %s: %w", id, ErrRecordNotFound)
}

// Write implements storage.ProjectionReadWriter.
func (f *FooDB) Write(ctx context.Context, val MeterReading) error {
	f.cache[val.MeterID] = val
	return nil
}

var _ storage.TableProjectionReadWriter[MeterReading] = &FooDB{}

func NewChargersDB() *FooDB {
	return &FooDB{make(map[string]MeterReading)}
}
