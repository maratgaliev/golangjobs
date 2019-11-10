package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/maratgaliev/golangjobs/config"
	"github.com/maratgaliev/golangjobs/model"
)

type Store interface {
	Close() error
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error

	GetJob(tx *sql.Tx, id int) (*model.Job, error)
	GetJobs(tx *sql.Tx) ([]*model.Job, error)
	CreateJob(tx *sql.Tx, job *model.Job) (*int, error)
	UpdateJob(tx *sql.Tx, job *model.Job) error
	DeleteJob(tx *sql.Tx, id int) error
}

type StoreContext struct {
	db *sql.DB
}

func NewStore(conf *config.Config) (Store, error) {
	storeConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Store.Host, conf.Store.Port, conf.Store.User, conf.Store.Password, conf.Store.Dbname,
	)
	db, err := sql.Open("postgres", storeConfig)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &StoreContext{db}, nil
}

func (sc *StoreContext) Close() error {
	return sc.db.Close()
}

func (sc *StoreContext) Begin() (*sql.Tx, error) {
	return sc.db.Begin()
}

func (sc *StoreContext) Commit(tx *sql.Tx) error {
	if tx == nil {
		return errors.New("tx is nil")
	}
	return tx.Commit()
}

func (sc *StoreContext) Rollback(tx *sql.Tx) error {
	if tx == nil {
		return errors.New("tx is nil")
	}
	return tx.Rollback()
}

func (sc *StoreContext) GetJob(tx *sql.Tx, id int) (*model.Job, error) {
	var query = "SELECT * FROM jobs WHERE id= $1;"
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRow(query, id)
	} else {
		row = sc.db.QueryRow(query, id)
	}
	job := &model.Job{}
	if err := row.Scan(&job.Id, &job.Title, &job.Description, &job.City, &job.Email, &job.Company, &job.Phone, &job.ContactName, &job.CurrencyType, &job.EmploymentType, &job.Salary, &job.IsApproved, &job.Website, &job.CreatedAt); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return job, nil
}

func (sc *StoreContext) GetJobs(tx *sql.Tx) ([]*model.Job, error) {
	query := "SELECT * FROM jobs order by id desc;"
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query(query)
	} else {
		rows, err = sc.db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobs []*model.Job
	for rows.Next() {
		job := &model.Job{}
		if err := rows.Scan(&job.Id, &job.Title, &job.Description, &job.City, &job.Email, &job.Company, &job.Phone, &job.ContactName, &job.CurrencyType, &job.EmploymentType, &job.Salary, &job.IsApproved, &job.Website, &job.CreatedAt); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (sc *StoreContext) CreateJob(tx *sql.Tx, job *model.Job) (*int, error) {
	var query = "INSERT INTO jobs(title, description, city, email, company, phone, contact_name, currency_type, employment_type, salary, is_approved, website, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;"
	var id int
	var err error
	dt := time.Now()
	if tx != nil {
		err = tx.QueryRow(query, job.Title, job.Description, job.City, job.Email, job.Company, job.Phone, job.ContactName, job.CurrencyType, job.EmploymentType, job.Salary, job.IsApproved, job.Website, dt).Scan(&id)
	} else {
		err = sc.db.QueryRow(query, job.Title, job.Description, job.City, job.Email, job.Company, job.Phone, job.ContactName, job.CurrencyType, job.EmploymentType, job.Salary, job.IsApproved, job.Website, dt).Scan(&id)
	}
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (sc *StoreContext) UpdateJob(tx *sql.Tx, job *model.Job) error {
	query := "UPDATE jobs SET title =$1 WHERE id = $2;"
	var res sql.Result
	var err error
	if tx != nil {
		res, err = tx.Exec(query, job.Title, job.Id)
	} else {
		res, err = sc.db.Exec(query, job.Title, job.Id)
	}
	if err != nil {
		return err
	}
	if a, err := res.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (sc *StoreContext) DeleteJob(tx *sql.Tx, id int) error {
	query := "DELETE FROM jobs WHERE id = $1;"
	var res sql.Result
	var err error
	if tx != nil {
		res, err = tx.Exec(query, id)
	} else {
		res, err = sc.db.Exec(query, id)
	}
	if err != nil {
		return err
	}
	if a, err := res.RowsAffected(); err != nil {
		return err
	} else if a == 0 {
		return sql.ErrNoRows
	}
	return nil
}
