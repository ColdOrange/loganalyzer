package loganalyzer

import "encoding/json"

type JsonError struct {
	Status string   `json:"status"`
	Errors []string `json:"errors"`
}

func jsonError(errs ...interface{}) []byte {
	result := &JsonError{
		Status: "failed",
	}
	for _, err := range errs {
		switch err.(type) {
		case string:
			result.Errors = append(result.Errors, err.(string))
		case error:
			result.Errors = append(result.Errors, err.(error).Error())
		}
	}
	data, err := json.Marshal(result)
	if err != nil {
		data = []byte(`{"status": "failed"}`)
	}
	return data
}

func jsonSuccess() []byte {
	return []byte(`{"status": "successful"}`)
}
