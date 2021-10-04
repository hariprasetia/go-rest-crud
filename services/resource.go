package services

import "go-rest-crud/utils"

type response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Error     string `json:"error,omitempty"`
	User interface{} `json:"user,omitempty"`
	Timestamp string `json:"timestamp"`
}

func (r *response) setResponse(s, msg, err string) {
	r.Status = s
	r.Message = msg
	r.Error = err
	r.Timestamp = utils.Timestamp()
}