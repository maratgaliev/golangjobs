package test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/maratgaliev/golangjobs/model"
	"github.com/maratgaliev/golangjobs/service"
	"github.com/maratgaliev/golangjobs/test/mock"
)

func TestJobService_GetJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockStore := mock.NewMockStore(mockCtrl)
	mockStore.EXPECT().GetJob(nil, 1).Return(nil, errors.New("test")).Times(1)
	mockStore.EXPECT().GetJob(nil, 2).Return(&model.Job{Id: 2, Title: "test"}, nil).Times(1)
	cs := service.NewJobService(mockStore)
	r, e := cs.GetJob(1)
	assert.NotNil(t, e)
	assert.Nil(t, r)
	r, e = cs.GetJob(2)
	assert.Nil(t, e)
	assert.NotNil(t, r)
}

func TestJobService_CreateJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mock.NewMockStore(mockCtrl)
	mockStore.EXPECT().Begin().Return(nil, errors.New("test")).Times(1)
	cs := service.NewJobService(mockStore)
	r, e := cs.CreateJob(&model.Job{Title: "Test"})
	assert.NotNil(t, e)
	assert.Nil(t, r)

	mockStore = mock.NewMockStore(mockCtrl)
	tx := new(sql.Tx)
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().CreateJob(tx, &model.Job{Title: "Test"}).Return(nil, errors.New("test")).Times(1)
	mockStore.EXPECT().Rollback(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	r, e = cs.CreateJob(&model.Job{Title: "Test"})
	assert.NotNil(t, e)
	assert.Nil(t, r)

	mockStore = mock.NewMockStore(mockCtrl)
	tx = new(sql.Tx)
	var id = 1
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().CreateJob(tx, &model.Job{Title: "Test"}).Return(&id, nil).Times(1)
	mockStore.EXPECT().Commit(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	r, e = cs.CreateJob(&model.Job{Title: "Test"})
	assert.Nil(t, e)
	assert.NotNil(t, r)
}

func TestJobService_UpdateJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mock.NewMockStore(mockCtrl)
	mockStore.EXPECT().Begin().Return(nil, errors.New("test")).Times(1)
	cs := service.NewJobService(mockStore)
	e := cs.UpdateJob(nil)
	assert.NotNil(t, e)

	mockStore = mock.NewMockStore(mockCtrl)
	tx := new(sql.Tx)
	cat := &model.Job{Id: 1, Title: "test"}
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().UpdateJob(tx, cat).Return(errors.New("test")).Times(1)
	mockStore.EXPECT().Rollback(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	e = cs.UpdateJob(cat)
	assert.NotNil(t, e)

	mockStore = mock.NewMockStore(mockCtrl)
	tx = new(sql.Tx)
	cat = &model.Job{Id: 1, Title: "test"}
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().UpdateJob(tx, cat).Return(nil).Times(1)
	mockStore.EXPECT().Commit(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	e = cs.UpdateJob(cat)
	assert.Nil(t, e)
}

func TestJobService_DeleteJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mock.NewMockStore(mockCtrl)
	mockStore.EXPECT().Begin().Return(nil, errors.New("test")).Times(1)
	cs := service.NewJobService(mockStore)
	e := cs.DeleteJob(1)
	assert.NotNil(t, e)

	mockStore = mock.NewMockStore(mockCtrl)
	tx := new(sql.Tx)
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().DeleteJob(tx, 1).Return(errors.New("test")).Times(1)
	mockStore.EXPECT().Rollback(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	e = cs.DeleteJob(1)
	assert.NotNil(t, e)

	mockStore = mock.NewMockStore(mockCtrl)
	tx = new(sql.Tx)
	mockStore.EXPECT().Begin().Return(tx, nil).Times(1)
	mockStore.EXPECT().DeleteJob(tx, 1).Return(nil).Times(1)
	mockStore.EXPECT().Commit(tx).Return(nil).Times(1)
	cs = service.NewJobService(mockStore)
	e = cs.DeleteJob(1)
	assert.Nil(t, e)
}
