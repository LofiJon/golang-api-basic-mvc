package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/models"
	request "todo-api/requests"
	"todo-api/services"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type TaskController struct {
	service   services.TaskService
	validator *validator.Validate
}

func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{
		service:   service,
		validator: validator.New(),
	}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body request.TaskRequest true "Task to create"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal error"
// @Router /tasks [post]
func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req request.TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	err := c.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := models.Task{
		Name: req.Name,
		Done: req.Done,
	}

	err = c.service.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// UpdateTask godoc
// @Summary Update a task by ID
// @Description Update a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body request.TaskRequest true "Task to update"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal error"
// @Router /tasks/{id} [put]
func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var req request.TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	err := c.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := models.Task{
		ID:   uint(id),
		Name: req.Name,
		Done: req.Done,
	}

	err = c.service.UpdateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description Delete a task by ID
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal error"
// @Router /tasks/{id} [delete]
func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteTask(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Internal error"
// @Router /tasks [get]
func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.service.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal error"
// @Router /tasks/{id} [get]
func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := c.service.GetTask(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}
