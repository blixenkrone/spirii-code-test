package chargers

import (
	"blixenkrone/spirii/storage"
	"context"
	"errors"
	"fmt"
	"sort"
	"time"
)

var (
	ErrRecordNotFound = errors.New("foo record not found")
)

type FooDB struct {
	cache map[string]MeterReading
}

type MeterReading struct {
	Timestamp       time.Time `json:"timestamp"`
	MeterID         string    `json:"meterID"`
	ConsumerID      string    `json:"consumerID"`
	MeterReadingVal int       `json:"meterReadingVal"`
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

func (f *FooDB) TopConsumers(ctx context.Context) ([]MeterReading, error) {

	if len(f.cache) == 0 {
		return nil, ErrRecordNotFound
	}

	out := make([]MeterReading, len(f.cache))
	idx := 0
	for _, v := range f.cache {
		out[idx] = MeterReading{
			Timestamp:       v.Timestamp,
			MeterID:         v.MeterID,
			ConsumerID:      v.ConsumerID,
			MeterReadingVal: v.MeterReadingVal,
		}
		idx++
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].MeterReadingVal > out[j].MeterReadingVal
	})

	return out, nil
}

var _ storage.TableProjectionReadWriter[MeterReading] = &FooDB{}

func NewChargersDB() *FooDB {
	return &FooDB{make(map[string]MeterReading)}
}
