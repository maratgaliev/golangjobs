package service

import (
	"github.com/maratgaliev/golangjobs/model"
	"github.com/maratgaliev/golangjobs/store"
)

type JobService interface {
	GetJob(id int) (*model.Job, error)
	GetJobs() ([]*model.Job, error)
	CreateJob(job *model.Job) (*int, error)
	UpdateJob(job *model.Job) error
	DeleteJob(id int) error
}

func NewJobService(store store.Store) JobService {
	return &JobServiceContext{store: store}
}

type JobServiceContext struct {
	store store.Store
}

func (csc *JobServiceContext) GetJob(id int) (*model.Job, error) {
	return csc.store.GetJob(nil, id)
}

func (csc *JobServiceContext) GetJobs() ([]*model.Job, error) {
	return csc.store.GetJobs(nil)
}

func (csc *JobServiceContext) CreateJob(job *model.Job) (*int, error) {
	tx, err := csc.store.Begin()
	if err != nil {
		return nil, err
	}
	cat, err := csc.store.CreateJob(tx, job)
	if err != nil {
		csc.store.Rollback(tx)
		return nil, err
	}
	if err = csc.store.Commit(tx); err != nil {
		return nil, err
	}
	return cat, nil
}

func (csc *JobServiceContext) UpdateJob(job *model.Job) error {
	tx, err := csc.store.Begin()
	if err != nil {
		return err
	}
	err = csc.store.UpdateJob(tx, job)
	if err != nil {
		csc.store.Rollback(tx)
		return err
	}
	if err = csc.store.Commit(tx); err != nil {
		return err
	}
	return nil
}

func (csc *JobServiceContext) DeleteJob(id int) error {
	tx, err := csc.store.Begin()
	if err != nil {
		return err
	}
	err = csc.store.DeleteJob(tx, id)
	if err != nil {
		csc.store.Rollback(tx)
		return err
	}
	if err = csc.store.Commit(tx); err != nil {
		return err
	}
	return nil
}
