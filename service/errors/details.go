package errors

type Details struct {
	Status  int
	Key     string
	Message string
}

type Available struct {
	Details map[int]Details
}

func (a Available) Result(key int, reasons ...string) *Result {
	details := a.Details[key]

	return &Result{
		Status:  details.Status,
		Key:     details.Key,
		Message: details.Message,
		Reasons: reasons,
	}
}
