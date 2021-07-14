package data

type SuccessfulRequest struct {
	Done bool `json: "done"`
	//On time.Time
}
//TODO: See how to create a timestamp for every request
const (
	layoutISO = "2000-01-02"
)
func NewSuccessfulRequest() *SuccessfulRequest{
	return &SuccessfulRequest{
		Done: true,
		//On: time.Parse(layoutISO,time.Now().String()),
	}
}
