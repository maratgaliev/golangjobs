package test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/maratgaliev/golangjobs/config"
	"github.com/maratgaliev/golangjobs/model"
	"github.com/maratgaliev/golangjobs/store"
	"github.com/stretchr/testify/assert"
)

var st store.Store

func init() {
	conf, _ := config.NewConfig("../config/config.yaml")
	st, _ = store.NewStore(conf)
}

func TestStore_CreateJob(t *testing.T) {
	tx, _ := st.Begin()
	defer st.Rollback(tx)
	id, err := st.CreateJob(tx, &model.Job{Title: "text"})
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestStore_GetJob(t *testing.T) {
	c, err := st.GetJob(nil, -1)
	assert.Nil(t, err)
	assert.Nil(t, c)
	tx, _ := st.Begin()
	defer st.Rollback(tx)
	id, _ := st.CreateJob(tx, &model.Job{Title: "text"})
	cat, _ := st.GetJob(tx, *id)
	assert.Equal(t, *id, cat.Id)
}

func TestStore_GetJobs(t *testing.T) {
	_, err := st.GetJobs(nil)
	assert.NoError(t, err)
	tx, _ := st.Begin()
	defer st.Rollback(tx)
	st.CreateJob(tx, &model.Job{Title: "text", IsApproved: true})
	res, _ := st.GetJobs(tx)
	assert.NotEmpty(t, res)
}

func TestStore_UpdateJob(t *testing.T) {
	tx, _ := st.Begin()
	defer st.Rollback(tx)
	id, _ := st.CreateJob(tx, &model.Job{Title: "text"})
	cat, _ := st.GetJob(tx, *id)
	cat.Title = "text2"
	st.UpdateJob(tx, cat)
	cat2, _ := st.GetJob(tx, cat.Id)
	assert.Equal(t, cat2.Title, "text2")
}

func TestStore_DeleteJob(t *testing.T) {
	tx, _ := st.Begin()
	defer st.Rollback(tx)
	id, _ := st.CreateJob(tx, &model.Job{Title: "text"})
	st.DeleteJob(tx, *id)
	cat2, _ := st.GetJob(tx, *id)
	assert.Nil(t, cat2)
}
