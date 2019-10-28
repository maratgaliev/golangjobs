// github.com/maratgaliev/golangjobs.
// Golang API
// Version: 0.1.3
// Schemes: http
// Host: localhost
// BasePath: /api
// Consumes:
// - application/json
// Produces:
// - application/json
// Contact: kazanlug@gmail.com
// swagger:meta
package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/maratgaliev/golangjobs/config"
	"github.com/maratgaliev/golangjobs/model"
	"github.com/maratgaliev/golangjobs/service"
	"gopkg.in/go-playground/validator.v9"
)

type Api struct {
	Http     *echo.Echo
	conf     *config.Config
	js       service.JobService
	apiInfo  ApiInfo
	validate *validator.Validate
}

type ApiInfo struct {
	Address string
	MW      []string
	Routes  []string
}

func ApiRunner(conf *config.Config, js service.JobService) *Api {
	api := &Api{}
	api.validate = validator.New()
	api.conf = conf
	api.js = js
	api.Http = echo.New()
	api.Http.Logger.SetLevel(log.Lvl(conf.LogLevel))
	api.apiInfo.Address = ":" + strconv.Itoa(api.conf.Api.HttpPort)
	api.Http.HideBanner = true
	api.Http.Pre(middleware.RemoveTrailingSlash())
	if conf.Api.Logging {
		api.Http.Use(middleware.Logger())
		api.apiInfo.MW = append(api.apiInfo.MW, "Logger")
	}
	api.Http.GET("/", api.index)
	api.Http.Static("/spec", "spec")

	api.Http.GET("/api/jobs", api.getJobs)
	api.Http.GET("/api/jobs/:id", api.getJob)
	api.Http.POST("/api/jobs", api.createJob)
	api.Http.PUT("/api/jobs/:id", api.updateJob)
	api.Http.DELETE("/api/jobs/:id", api.deleteJob)

	for _, r := range api.Http.Routes() {
		api.apiInfo.Routes = append(api.apiInfo.Routes, fmt.Sprintf("%s %s", r.Path, r.Method))
	}
	return api
}

// Запустить api
func (api *Api) Start() error {
	return api.Http.Start(":" + strconv.Itoa(api.conf.Api.HttpPort))
}

// Инфо об api
func (api *Api) GetApiInfo() ApiInfo {
	return api.apiInfo
}

func (api *Api) index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/spec/api.json")
}

func (api *Api) spec(c echo.Context) error {
	return c.Inline("spec/api.json", "api.json")
}

// swagger:operation GET /jobs/{id} getJob
// ---
// description: Get job
// parameters:
// - name: id
//   in: path
//   description: job id
//   required: true
//   type: int
// responses:
//  '200':
//    schema:
//      $ref: '#/definitions/Job'
//  '400':
//     description: Bad request param `id`
//  '404':
//     description: Job `id`= not found
//
func (api *Api) getJob(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param `id`")
	}
	cat, err := api.js.GetJob(id)
	if err != nil {
		return err
	}
	if cat == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Job `id` = ", id, " not found")
	}
	return c.JSON(http.StatusOK, cat)
}

// swagger:operation GET /jobs getJobs
// ---
// description: Get jobs list
// responses:
//  '200':
//    schema:
//      type: array
//      items:
//        $ref: '#/definitions/Job'
//
func (api *Api) getJobs(c echo.Context) error {
	cats, err := api.js.GetJobs()
	if err != nil {
		return err
	}
	if cats == nil {
		cats = []*model.Job{}
	}
	return c.JSON(http.StatusOK, cats)
}

// swagger:operation POST /jobs createJob
// ---
// description: Create job
// parameters:
// - name: job
//   in: body
//   description: new job
//   required: true
//   schema:
//     $ref: '#/definitions/Job'
// responses:
//  '201':
//    schema:
//      $ref: '#/definitions/Job'
//  '400':
//     description: Bad request param
//
func (api *Api) createJob(c echo.Context) error {
	req := &model.Job{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param: "+err.Error())
		return err
	}
	if err := api.validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param: "+err.Error())
	}
	res, err := api.js.CreateJob(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]*int{"id": res})
}

// swagger:operation PUT /jobs/{id} updateJob
// ---
// description: Update job
// parameters:
// - name: id
//   in: path
//   description: job id
//   required: true
//   type: int
// - name: job
//   in: body
//   description: updated job
//   required: true
//   schema:
//     $ref: '#/definitions/Job'
// responses:
//  '204':
//     description: Job updated
//  '400':
//     description: Bad request param
//
func (api *Api) updateJob(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param `id`")
	}
	req := &model.Job{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param: "+err.Error())
	}
	if err := api.validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param: "+err.Error())
	}
	req.Id = id
	if err = api.js.UpdateJob(req); err != nil {
		if err != sql.ErrNoRows {
			return err
		} else {
			return echo.NewHTTPError(http.StatusNotFound, "Job `id` = ", id, " not found")
		}

	}
	return c.NoContent(http.StatusNoContent)
}

// swagger:operation DELETE /jobs/{id} deleteJob
// ---
// description: Delete job
// parameters:
// - name: id
//   in: path
//   description: job id
//   required: true
//   type: int
// responses:
//  '204':
//     description: Job deleted
//  '400':
//     description: Bad request param `id`
//  '404':
//     description: Job `id`= not found
//
func (api *Api) deleteJob(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request param `id`")
	}
	if err = api.js.DeleteJob(id); err != nil {
		if err != sql.ErrNoRows {
			return err
		} else {
			return echo.NewHTTPError(http.StatusNotFound, "Job `id` = ", id, " not found")
		}
	}

	return c.NoContent(http.StatusNoContent)
}
