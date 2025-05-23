package services

import (
	"deadshot/internal/models"
	"deadshot/internal/persistence"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type SqlliteDeadshotService struct {
	Repository *persistence.DBManager
}

func NewSqlliteDeadshotService(d *persistence.DBManager) *SqlliteDeadshotService {
	return &SqlliteDeadshotService{Repository: d}
}

func (s *SqlliteDeadshotService) InsertLog(log *models.LogModel) error {
	_, err := s.Repository.InsertLog(log)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlliteDeadshotService) GetAllLogs() ([]models.LogModel, error) {
	logs, err := s.Repository.GetAllLogs()

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (s *SqlliteDeadshotService) GetLogById(id int) (*models.LogModel, error) {
	log, err := s.Repository.GetLogByID(id)

	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *SqlliteDeadshotService) UpdateLog(log *models.LogModel) error {
	err := s.Repository.UpdateLog(log)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlliteDeadshotService) DeleteLog(id int) error {
	err := s.Repository.DeleteLog(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlliteDeadshotService) IncreaseRequestCount(id int) error {
	err := s.Repository.IncreaseReplayCount(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqlliteDeadshotService) ReplayLog(id int) error {
	log, geterr := s.Repository.GetLogByID(id)

	if geterr != nil {
		return geterr
	}

	httpClient := &http.Client{}

	req, reqerr := http.NewRequest(log.Method, log.URL, nil)
	if reqerr != nil {
		logrus.Error("Error creating request:" + reqerr.Error())
		return reqerr
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	var result map[string]interface{}
	maperr := json.Unmarshal([]byte(log.Headers), &result)

	if maperr != nil {
		logrus.Error(maperr)
		return maperr
	}

	for key, value := range result {
		req.Header.Set(key, fmt.Sprintf("%s", value))
	}
	resp, doerr := httpClient.Do(req)
	if doerr != nil {
		logrus.Error("Error sending request:" + doerr.Error())
		return doerr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		incerr := s.IncreaseRequestCount(id)
		if incerr != nil {
			logrus.Error("Error increasing request count:" + incerr.Error())
			return incerr
		}
	}

	return nil
}
