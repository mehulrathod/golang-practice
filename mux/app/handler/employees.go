package handler

import (
	"HirebasePractice/mux/model"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func GetAllEmployees(db *gorm.DB,w http.ResponseWriter,r *http.Request)  {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB,w http.ResponseWriter,r *http.Request)  {
	employee := model.Employee{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee);
	err != nil{
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error;
	err != nil{
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}