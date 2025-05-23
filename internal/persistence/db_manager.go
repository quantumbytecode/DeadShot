package persistence

import (
	"database/sql"
	"deadshot/internal/models"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

// SetupDB initializes the database connection and creates tables
func (m *DBManager) SetupDB() {

	createTable := `
	CREATE TABLE IF NOT EXISTS records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		method TEXT,
		url TEXT,
		headers TEXT,
		query_params TEXT,
		body TEXT,
		received_at DATETIME,
		status_code INTEGER,
		response_headers TEXT,
		response_body TEXT,
		tags TEXT,
		source TEXT,
		replayed BOOLEAN DEFAULT 0,
		error TEXT
	);`

	if _, err := m.db.Exec(createTable); err != nil {
		logrus.Fatal("Failed to create table:", err)
	}

	logrus.Info("Database initialized successfully")
}

// InsertLog creates a new record
func (m *DBManager) InsertLog(log *models.LogModel) (int64, error) {
	headers, _ := json.Marshal(log.Headers)
	queryParams, _ := json.Marshal(log.QueryParams)
	responseHeaders, _ := json.Marshal(log.ResponseHeaders)
	tags, _ := json.Marshal(log.Tags)

	res, err := m.db.Exec(
		`INSERT INTO records (
			method, url, headers, query_params, body, received_at,
			status_code, response_headers, response_body, tags,
			source, replayed, error
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		log.Method,
		log.URL,
		string(headers),
		string(queryParams),
		log.Body,
		time.Now().Format(time.RFC3339),
		log.StatusCode,
		string(responseHeaders),
		log.ResponseBody,
		string(tags),
		log.Source,
		log.Replayed,
		log.Error,
	)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetLogByID retrieves a single record by ID
func (m *DBManager) GetLogByID(id int) (*models.LogModel, error) {
	var log models.LogModel
	var headers, queryParams, responseHeaders, tags string
	var ReceivedAt string

	err := m.db.QueryRow(`
		SELECT 
			id, method, url, headers, query_params, body, received_at,
			status_code, response_headers, response_body, tags,
			source, replayed, error
		FROM records WHERE id = ?`, id).Scan(
		&log.ID,
		&log.Method,
		&log.URL,
		&headers,
		&queryParams,
		&log.Body,
		&ReceivedAt,
		&log.StatusCode,
		&responseHeaders,
		&log.ResponseBody,
		&tags,
		&log.Source,
		&log.Replayed,
		&log.Error,
	)

	if err != nil {
		return nil, err
	}

	// Parse JSON fields
	json.Unmarshal([]byte(headers), &log.Headers)
	json.Unmarshal([]byte(queryParams), &log.QueryParams)
	json.Unmarshal([]byte(responseHeaders), &log.ResponseHeaders)
	json.Unmarshal([]byte(tags), &log.Tags)

	// Parse timestamp
	log.ReceivedAt, _ = time.Parse(time.RFC3339, ReceivedAt)

	return &log, nil
}

// UpdateLog modifies an existing record
func (m *DBManager) UpdateLog(log *models.LogModel) error {
	headers, _ := json.Marshal(log.Headers)
	queryParams, _ := json.Marshal(log.QueryParams)
	responseHeaders, _ := json.Marshal(log.ResponseHeaders)
	tags, _ := json.Marshal(log.Tags)

	_, err := m.db.Exec(`
		UPDATE records SET
			method = ?,
			url = ?,
			headers = ?,
			query_params = ?,
			body = ?,
			status_code = ?,
			response_headers = ?,
			response_body = ?,
			tags = ?,
			source = ?,
			replayed = ?,
			error = ?
		WHERE id = ?`,
		log.Method,
		log.URL,
		string(headers),
		string(queryParams),
		log.Body,
		log.StatusCode,
		string(responseHeaders),
		log.ResponseBody,
		string(tags),
		log.Source,
		log.Replayed,
		log.Error,
		log.ID,
	)

	return err
}

// DeleteLog removes a record by ID
func (m *DBManager) DeleteLog(id int) error {
	_, err := m.db.Exec("DELETE FROM records WHERE id = ?", id)
	return err
}

// GetAllLogs returns all records with pagination
func (m *DBManager) GetAllLogs() ([]models.LogModel, error) {
	rows, err := m.db.Query(`
		SELECT 
			id, method, url, headers, query_params, body, received_at,
			status_code, response_headers, response_body, tags,
			source, replayed, error
		FROM records
		ORDER BY received_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.LogModel
	for rows.Next() {
		var log models.LogModel
		var headers, queryParams, responseHeaders, tags string
		var ReceivedAt string

		err := rows.Scan(
			&log.ID,
			&log.Method,
			&log.URL,
			&headers,
			&queryParams,
			&log.Body,
			&ReceivedAt,
			&log.StatusCode,
			&responseHeaders,
			&log.ResponseBody,
			&tags,
			&log.Source,
			&log.Replayed,
			&log.Error,
		)
		if err != nil {
			return nil, err
		}

		// Parse JSON fields
		json.Unmarshal([]byte(headers), &log.Headers)
		json.Unmarshal([]byte(queryParams), &log.QueryParams)
		json.Unmarshal([]byte(responseHeaders), &log.ResponseHeaders)
		json.Unmarshal([]byte(tags), &log.Tags)

		// Parse timestamp
		log.ReceivedAt, _ = time.Parse(time.RFC3339, ReceivedAt)

		logs = append(logs, log)
	}

	return logs, nil
}

// Close cleans up the database connection
func (m *DBManager) Close() error {
	return m.db.Close()
}

func (m *DBManager) IncreaseReplayCount(id int) error {
	_, err := m.db.Exec(`
		UPDATE records SET
			replayed_count = replayed_count + 1
		WHERE id = ?`,
		id,
	)

	if err != nil {
		logrus.Error("Failed to increase request count:", err)
		return err
	}

	return nil
}
