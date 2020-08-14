package main

import (
	"net/http"

	"github.com/archit-p/MicroserviceTemplate/pkg/dto"

	"github.com/jinzhu/copier"
)

// createSample godoc
// @Summary Create a Sample in the database
// @Description Create a Sample in the database
// @Tags samples
// @Accept json
// @Produce json
// @Param sample body dto.Sample true "Create sample"
// @Success 200 ID string
// @Router /sample/create [post]
func (app *application) createSample(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	var s dto.Sample
	err := s.FromJSON(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.samples.Insert(s.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Created ", id)
	w.Write([]byte("Created document"))
}

// getSample godoc
// @Summary Retrive a Sample from the database
// @Description Retrive a Sample from the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [get]
func (app *application) getSample(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	snip, err := app.samples.Get(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var s dto.Sample
	copier.Copy(&s, &snip)

	err = s.ToJSON(w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Read ", id)
}

// updateSample godoc
// @Summary Update a Sample in the database
// @Description Update a Sample in the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Param sample body dto.Sample true "Update sample"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [put]
func (app *application) updateSample(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get(":id")

	var s dto.Sample
	err := s.FromJSON(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, err := app.samples.Update(id, s.Content)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infoLog.Printf("Updated : %d : %s", res, id)
	w.Write([]byte("Updated document"))
}

// deleteSample godoc
// @Summary Delete a Sample in the database
// @Description Delete a Sample in the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [delete]
func (app *application) deleteSample(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	res, err := app.samples.Delete(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infoLog.Printf("Deleted : %d : %s", res, id)
	w.Write([]byte("Deleted document"))
}
