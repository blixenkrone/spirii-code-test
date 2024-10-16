package http

import (
	"blixenkrone/spirii/internal/chargers"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetCourse(t *testing.T) {
	t.Run("API can write and get a chargers record", func(t *testing.T) {
		chargersID := "1"
		chargersDB := chargers.NewChargersDB()
		t.Run("Responds 404 for no record", func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			req = mux.SetURLVars(req, map[string]string{"id": chargersID})

			s := Server{
				logger: logrus.New(),
				memDB:  chargersDB,
			}
			s.getMeterDataV1()(rr, req)
			assert.Equal(t, 404, rr.Code)

		})
		t.Run("Write chargers record", func(t *testing.T) {
			// rr := httptest.NewRecorder()
			// body := strings.NewReader(`
			// {
			// 	"id": "1",
			// 	"value": "hello, world!"
			// }`)
			// req := httptest.NewRequest("POST", "/", body)

			// s := Server{
			// 	logger: logrus.New(),
			// 	memDB:  chargersDB,
			// }
			// s.postchargersV1()(rr, req)
			// assert.Equal(t, 202, rr.Code)

		})
		t.Run("Read chargers record", func(t *testing.T) {
			// rr := httptest.NewRecorder()
			// req := httptest.NewRequest("GET", "/chargers/{id}", nil)
			// req = mux.SetURLVars(req, map[string]string{"id": chargersID})

			// s := Server{
			// 	logger: logrus.New(),
			// 	memDB:  chargersDB,
			// }
			// s.getchargersV1()(rr, req)
			// assert.Equal(t, 200, rr.Code)
		})

	})
}
