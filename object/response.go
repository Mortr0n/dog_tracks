package object

type ErrorResponse struct {
	Status int
	Error  string
}

type SuccessResponse struct {
	Status   int
	Response any
}
