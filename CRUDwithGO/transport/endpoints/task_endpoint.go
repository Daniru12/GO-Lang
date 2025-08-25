package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"patricego/services"
	"patricego/usecases/domain"
)

type TaskData struct {
	ID          int64  `json:"id,omitempty"`
	ResourceID  string `json:"resource_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
	Status      string `json:"status"`
}


type TaskHandler struct {
	TaskService *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: service}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, message string, resourceID string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"code":       statusCode,
		"message":    message,
		"resourceId": resourceID,
		"data":       data,
	}

	json.NewEncoder(w).Encode(response)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string, resourceID string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"code":       statusCode,
		"message":    message,
		"resourceId": resourceID,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid request", "")
		return
	}
	if task.Name == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Task name is required", "")
		return
	}

	loc, _ := time.LoadLocation("Asia/Colombo")
	now := time.Now().In(loc)

	task.CreatedTime = now
	task.UpdatedTime = now
	task.Status = "A"
	task.ResourceID = uuid.New().String()

	id, err := h.TaskService.CreateTask(r.Context(), task)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to create task", task.ResourceID)
		return
	}
	task.Id = id

	writeJSONResponse(w, http.StatusCreated, "Task created successfully", task.ResourceID, map[string]bool{
		"result": true,
	})
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.TaskService.GetAllTasks(r.Context())
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve tasks", "")
		return
	}

	loc, _ := time.LoadLocation("Asia/Colombo")
	var taskList []TaskData
	for _, task := range tasks {
		taskList = append(taskList, TaskData{
			Name:        task.Name,
			Description: task.Description,
			ResourceID:  task.ResourceID,
			CreatedTime: task.CreatedTime.In(loc).Format("2006-01-02 15:04:05"),
			UpdatedTime: task.UpdatedTime.In(loc).Format("2006-01-02 15:04:05"),
			Status:      task.Status,
		})
	}

	response := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Tasks retrieved successfully",
		"data":    taskList,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetTaskByResourceID(w http.ResponseWriter, r *http.Request) {
	resourceID := mux.Vars(r)["resource_id"]
	task, err := h.TaskService.GetTaskByResourceID(r.Context(), resourceID)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "Task not found", resourceID)
		return
	}

	loc, _ := time.LoadLocation("Asia/Colombo")
	data := TaskData{
		Name:        task.Name,
		Description: task.Description,
		ResourceID:  task.ResourceID,
		CreatedTime: task.CreatedTime.In(loc).Format("2006-01-02 15:04:05"),
		UpdatedTime: task.UpdatedTime.In(loc).Format("2006-01-02 15:04:05"),
		Status:      task.Status,
	}

	response := map[string]interface{}{
		"code":       http.StatusOK,
		"message":    "Task retrieved successfully",
		"resourceId": resourceID, 
		"data":       data,
	}

	json.NewEncoder(w).Encode(response)
}



func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	resourceID := mux.Vars(r)["resource_id"]
	var updateData map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid request body", resourceID)
		return
	}

	task, err := h.TaskService.GetTaskByResourceID(r.Context(), resourceID)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "Task not found", resourceID)
		return
	}

	originalStatus := task.Status

	if name, ok := updateData["name"].(string); ok {
		task.Name = name
	}
	if desc, ok := updateData["description"].(string); ok {
		task.Description = desc
	}
	if status, ok := updateData["status"].(string); ok {
		task.Status = status
	}

	loc, _ := time.LoadLocation("Asia/Colombo")
	task.UpdatedTime = time.Now().In(loc)
	task.ResourceID = resourceID

	if err := h.TaskService.UpdateTask(r.Context(), task); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to update task", resourceID)
		return
	}

	message := "Task updated"
	if originalStatus == "A" && task.Status == "D" {
		message = "Task deleted successfully"
	}

	
	response := map[string]interface{}{
		"code":    http.StatusOK,
		"data": map[string]interface{}{
			"result": true,
				 
		},
		"message": message,
		"resourceId": resourceID, 
		
		 
	}

	json.NewEncoder(w).Encode(response)
}


 
func (h *TaskHandler) TaskDelete(w http.ResponseWriter, r *http.Request) {
	resourceID := mux.Vars(r)["resource_id"]

	err := h.TaskService.UpdateTaskStatus(r.Context(), resourceID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to delete task", resourceID)
		return
	}
	writeJSONResponse(w, http.StatusOK, "Task deleted", resourceID, map[string]bool{
		"result": true,
	})
}


