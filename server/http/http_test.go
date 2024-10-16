package http

import (
	"blixenkrone/spirii/internal/chargers"
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetCourse(t *testing.T) {
	t.Run("API can write and get a chargers record", func(t *testing.T) {
		ctx := context.TODO()
		chargerID := "1"
		chargersDB := chargers.NewChargersDB()
		t.Run("Responds 404 for no record", func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			req = mux.SetURLVars(req, map[string]string{"id": chargerID})

			s := Server{
				logger: logrus.New(),
				memDB:  chargersDB,
			}
			s.getMeterDataV1()(rr, req)
			assert.Equal(t, 404, rr.Code)

		})
		t.Run("Write chargers record", func(t *testing.T) {
			err := chargersDB.Write(ctx, chargers.MeterReading{
				Timestamp:       time.Now(),
				MeterID:         chargerID,
				ConsumerID:      "12345",
				MeterReadingVal: 16,
			})
			assert.NoError(t, err)
		})
		t.Run("Read chargers record", func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/chargers/{id}", nil)
			req = mux.SetURLVars(req, map[string]string{"id": chargerID})

			s := Server{
				logger: logrus.New(),
				memDB:  chargersDB,
			}
			s.getMeterDataV1()(rr, req)
			assert.Equal(t, 200, rr.Code)
		})

	})
}
