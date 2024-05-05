package errors

func ToInternalServerError(e error) *Result {
	return Available{Details: DefaultErrors}.
		Result(InternalServerError, e.Error())
}
