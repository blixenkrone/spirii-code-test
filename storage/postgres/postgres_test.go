package postgres

import (
	"testing"
)

func TestMigrations(t *testing.T) {
	// t.Run("with docker ressource", func(t *testing.T) {
	// 	a := assert.New(t)
	// 	pool, err := docker.NewPool()
	// 	a.NoError(err)
	// 	r, err := pool.Postgres("testdb")
	// 	a.NoError(err)
	// 	defer r.Teardown()

	// 	db := NewFromConn(r.Container())
	// 	err = db.RunMigrations("./migrations")
	// 	a.NoError(err)
	// })
}
