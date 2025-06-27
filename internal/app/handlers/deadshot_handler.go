package handlers

import (
	"deadshot/internal/app/services"
	"deadshot/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type DeadshotHandler struct {
	Service *services.SqlliteDeadshotService
}

func NewDeadshotHandler(s *services.SqlliteDeadshotService) *DeadshotHandler {
	return &DeadshotHandler{
		Service: s,
	}
}

func (d *DeadshotHandler) CaptureLog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "aspplication/json")

	body, err := io.ReadAll(req.Body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not read body, check it again",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	var log models.LogModel
	decodErr := json.Unmarshal(body, &log)

	if decodErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		logrus.Error(decodErr)
		response := models.DeadShotResponse{
			Message: "Could not parse body, check it again",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	inErr := d.Service.InsertLog(&log)
	if inErr != nil {
		logrus.Error(inErr)
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not insert log",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := models.DeadShotResponse{
		Message: "200",
		Data:    nil,
	}
	resp, _ := json.Marshal(response)
	res.Write(resp)
	logrus.Info("Log captured")
}

func (d *DeadshotHandler) GetAllLogs(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "aspplication/json")

	logs, err := d.Service.GetAllLogs()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not get logs",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := models.DeadShotResponse{
		Message: "200",
		Data:    logs,
	}
	resp, _ := json.Marshal(response)
	res.Write(resp)
	logrus.Info("Logs retrieved")
}

func (d *DeadshotHandler) GetLogByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	id := req.URL.Query().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}

	log, err := d.Service.GetLogById(i)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(models.DeadShotResponse{
			Message: "Could not get log",
			Data:    nil,
		})
		return
	}

	json.NewEncoder(res).Encode(models.DeadShotResponse{
		Message: "Log found",
		Data:    log,
	})

	logrus.Info("Log retrieved")
}

func (d *DeadshotHandler) DeleteLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "aspplication/json")

	id := req.URL.Query().Get("id")

	i, err := strconv.Atoi(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not parse id",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	err = d.Service.DeleteLog(i)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not delete log",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := models.DeadShotResponse{
		Message: "200",
		Data:    nil,
	}
	resp, _ := json.Marshal(response)
	res.Write(resp)
	logrus.Info("Log deleted")
}

func (d *DeadshotHandler) UpdateLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "aspplication/json")

	body, err := io.ReadAll(req.Body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not read body, check it again",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
	}

	var log models.LogModel
	decodErr := json.Unmarshal(body, &log)

	if decodErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		logrus.Error(decodErr)
		response := models.DeadShotResponse{
			Message: "Could not parse body, check it again",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
	}

	inErr := d.Service.UpdateLog(&log)
	if inErr != nil {
		logrus.Error(inErr)
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not update log",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := models.DeadShotResponse{
		Message: "200",
		Data:    nil,
	}
	resp, _ := json.Marshal(response)
	res.Write(resp)
	logrus.Info("Log updated")
}

func (d *DeadshotHandler) ReplayLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "aspplication/json")

	id := req.URL.Query().Get("id")

	i, err := strconv.Atoi(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not parse id",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	replayErr := d.Service.ReplayLog(i)

	if replayErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := models.DeadShotResponse{
			Message: "Could not replay log",
			Data:    nil,
		}
		e, _ := json.Marshal(response)
		res.Write(e)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := models.DeadShotResponse{
		Message: "200",
		Data:    nil,
	}
	resp, _ := json.Marshal(response)
	res.Write(resp)
	logrus.Info("Request replayed")
}
