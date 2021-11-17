package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang_crud/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var dbconn *gorm.DB

type Response struct {
	Data    []models.Task `json:"data"`
	Message string        `json:"message"`
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks = models.GetTasks()
	var resp Response
	err := dbconn.Find(&tasks).Error
	if err == nil {
		log.Println(tasks)
		resp.Data = tasks
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	var task = models.GetTask()
	err := dbconn.Where("id = ?", id).Find(&task).Error
	if err == nil {
		log.Println(task)
		resp.Data = append(resp.Data, task)
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	var task = models.GetTask()
	_ = json.NewDecoder(r.Body).Decode(&task)
	log.Println(task)

	err := dbconn.Create(&task).Error
	if err != nil {
		http.Error(w, "Error Creating Record", 400)
		return
	}
	resp.Message = "CREATED"
	json.NewEncoder(w).Encode(resp)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var task = models.GetTask()
	_ = json.NewDecoder(r.Body).Decode(&task)
	id, _ := strconv.Atoi(params["id"])

	err := dbconn.Model(&task).Where("id = ?", id).Update(&task).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp.Message = "UPDATED"
	json.NewEncoder(w).Encode(resp)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var task = models.GetTask()
	err := dbconn.Delete(&task, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	resp.Message = "DELETED"
	json.NewEncoder(w).Encode(resp)
}

func SetDB(db *gorm.DB) {
	dbconn = db
	var task = models.GetTask()
	dbconn.AutoMigrate(&task)
}
