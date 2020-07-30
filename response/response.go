package response

import "time"

type RetrieveAuthPayload struct {
	StatusCode string `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination interface{} `json:"pagination"`
		Auths      []struct {
			ID        string `json:"id"`
			Validated bool   `json:"validated"`
			Record    struct {
				ID   string `json:"_id"`
				Bank struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"bank"`
				Env      string `json:"env"`
				Owner    string `json:"owner"`
				Customer struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"customer"`
			} `json:"record"`
			CreatedAt   time.Time `json:"created_at"`
			LastUpdated time.Time `json:"last_updated"`
		} `json:"auths"`
	} `json:"data"`
}
