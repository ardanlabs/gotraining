package model

type StartRequest struct {
	Image string
	Count int
}

type StartResponse struct {
	ID  string
	Err error
}
