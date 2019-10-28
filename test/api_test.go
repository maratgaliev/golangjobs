package test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/maratgaliev/golangjobs/api"
	"github.com/maratgaliev/golangjobs/config"
	"github.com/maratgaliev/golangjobs/model"
	"github.com/maratgaliev/golangjobs/test/mock"
)

func TestApi_GetJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	conf := &config.Config{LogLevel: 5}
	cs := mock.NewMockJobService(mockCtrl)
	api := api.ApiRunner(conf, cs)
	req := httptest.NewRequest(echo.GET, "/api/jobs/2", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 404
	rec := httptest.NewRecorder()
	cs.EXPECT().GetJob(2).Return(nil, nil).Times(1)
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	// 200
	cat := &model.Job{Id: 2, Title: "Title2"}
	cs.EXPECT().GetJob(2).Return(cat, nil).Times(1)
	rec = httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	res, _ := json.Marshal(cat)
	assert.Equal(t, rec.Body.String(), string(res))
}

func TestApi_CreateJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	conf := &config.Config{LogLevel: 5}
	cs := mock.NewMockJobService(mockCtrl)
	api := api.ApiRunner(conf, cs)
	// 400
	catJSON := `{"title": ""}`
	req := httptest.NewRequest(echo.POST, "/api/jobs/", strings.NewReader(catJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 200
	catJSON = `{"title": "golang title", "description": "golang descr", "city": "kazan", "email": "hr@example.com", "company": "Company", "phone": "4167702450", "salary": 1000000}`
	id := 2
	req = httptest.NewRequest(echo.POST, "/api/jobs/", strings.NewReader(catJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cs.EXPECT().CreateJob(gomock.Any()).Return(&id, nil).Times(1)
	rec = httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var res map[string]int
	json.Unmarshal(rec.Body.Bytes(), &res)
	assert.Equal(t, res["id"], 2)
}

func TestApi_UpdateJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	conf := &config.Config{LogLevel: 5}
	cs := mock.NewMockJobService(mockCtrl)
	api := api.ApiRunner(conf, cs)
	// 400
	catJSON := `{"title": ""}`
	req := httptest.NewRequest(echo.PUT, "/api/jobs/2", strings.NewReader(catJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 404
	catJSON = `{"title": "golang title", "description": "golang descr", "city": "kazan", "email": "hr@example.com", "company": "Company", "phone": "4167702450", "salary": 1000000}`
	req = httptest.NewRequest(echo.PUT, "/api/jobs/1", strings.NewReader(catJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cs.EXPECT().UpdateJob(gomock.Any()).Return(sql.ErrNoRows).Times(1)
	rec = httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	// 204
	catJSON = `{"title": "golang title", "description": "golang descr", "city": "kazan", "email": "hr@example.com", "company": "Company", "phone": "4167702450", "salary": 1000000}`
	req = httptest.NewRequest(echo.PUT, "/api/jobs/2", strings.NewReader(catJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cs.EXPECT().UpdateJob(gomock.Any()).Return(nil).Times(1)
	rec = httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestApi_DeleteJob(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	conf := &config.Config{LogLevel: 5}
	cs := mock.NewMockJobService(mockCtrl)
	api := api.ApiRunner(conf, cs)
	// 404
	req := httptest.NewRequest(echo.DELETE, "/api/jobs/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cs.EXPECT().DeleteJob(gomock.Any()).Return(sql.ErrNoRows).Times(1)
	rec := httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	// 201
	req = httptest.NewRequest(echo.DELETE, "/api/jobs/2", nil)
	cs.EXPECT().DeleteJob(gomock.Any()).Return(nil).Times(1)
	rec = httptest.NewRecorder()
	api.Http.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
