package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KevinNguyen1703/FrameStoreGolangServer/pkg/models"
	"github.com/KevinNguyen1703/FrameStoreGolangServer/pkg/utils"
	"github.com/gorilla/mux"
)

var NewFrame models.Frame

func GetFrame(w http.ResponseWriter, r *http.Request) {
	newFrames := models.GetAllFrames()
	res, _ := json.Marshal(newFrames)
	w.Header().Set("Content-Type", "pkgaplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFrameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	frameId := vars["frameID"]
	ID, err := strconv.ParseInt(frameId, 10, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	frameDetails, _ := models.GetFrameById(ID)
	res, _ := json.Marshal(frameDetails)
	w.Header().Set("Content-Type", "dkgaplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFrame(w http.ResponseWriter, r *http.Request) {
	CreateFrame := &models.Frame{}
	utils.ParseBody(r, CreateFrame)
	b := CreateFrame.CreateFrame()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	frameId := vars["frameID"]
	ID, err := strconv.ParseInt(frameId, 10, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	frame := models.DeleteFrame(ID)
	res, _ := json.Marshal(frame)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	var index int64 = 1
	for {
		frame, _ := models.GetFrameById(index)
		if frame.Data == "" {
			break
		}
		models.DeleteFrame(index)
		index++
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateFrame(w http.ResponseWriter, r *http.Request) {
	var updateFrame = &models.Frame{}
	utils.ParseBody(r, updateFrame)
	vars := mux.Vars(r)
	frameId := vars["frameID"]
	ID, err := strconv.ParseInt(frameId, 10, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	frameDetails, db := models.GetFrameById(ID)
	if updateFrame.Video != "" {
		frameDetails.Video = updateFrame.Video
	}
	if updateFrame.Data != "" {
		frameDetails.Data = updateFrame.Data
	}
	if updateFrame.Description != "" {
		frameDetails.Description = updateFrame.Description
	}
	db.Save(&frameDetails)
	res, _ := json.Marshal(frameDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
