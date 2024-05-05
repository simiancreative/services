package errors

type Result struct {
	Status  int      `json:"status"`
	Key     string   `json:"key"`
	Message string   `json:"message"`
	Reasons []string `json:"reasons"`
}

func (m *Result) Error() string {
	return m.Message
}
