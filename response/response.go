package response

import "time"

type RetrieveAuthPayload []struct {
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

type AuthByIDPayload struct {
	StatusCode string `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int    `json:"totalDocs"`
			Limit         int    `json:"limit"`
			HasPrevPage   bool   `json:"hasPrevPage"`
			HasNextPage   bool   `json:"hasNextPage"`
			Page          int    `json:"page"`
			TotalPages    int    `json:"totalPages"`
			PagingCounter int    `json:"pagingCounter"`
			PrevPage      string `json:"prevPage"`
			NextPage      string `json:"nextPage"`
		} `json:"pagination"`
		Auths []struct {
			ID     string `json:"_id"`
			Record string `json:"record"`
			V      int    `json:"__v"`
			Bank   struct {
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Icon   string `json:"icon"`
				Status string `json:"status"`
			} `json:"bank"`
			CreatedAt time.Time `json:"created_at"`
			Customer  struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Env         string    `json:"env"`
			LastUpdated time.Time `json:"last_updated"`
			Owner       string    `json:"owner"`
			Validated   bool      `json:"validated"`
		} `json:"auths"`
	} `json:"data"`
}

type RetrieveBalancePayload []struct {
	StatusCode string `json:"statusCode"`
	ID         string `json:"id"`
	Customer   string `json:"customer"`
	Account    struct {
		Owner []struct {
		} `json:"owner"`
		Record []struct {
		} `json:"record"`
		ID   string `json:"_id"`
		Name string `json:"name"`
		Bank struct {
		} `json:"bank"`
		Customer struct {
		} `json:"customer"`
	} `json:"account"`
	LedgerBalance    float64   `json:"ledger_balance"`
	AvailableBalance int       `json:"available_balance"`
	Currency         string    `json:"currency"`
	Connected        []string  `json:"connected"`
	Env              string    `json:"env"`
	CreatedAt        time.Time `json:"created_at"`
	LastUpdated      time.Time `json:"last_updated"`
}

// id = "5f03b2d5af4681496afdd84f"

type BalanceByIDPayload struct {
	StatusCode string `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int    `json:"totalDocs"`
			Limit         int    `json:"limit"`
			HasPrevPage   bool   `json:"hasPrevPage"`
			HasNextPage   bool   `json:"hasNextPage"`
			Page          int    `json:"page"`
			TotalPages    int    `json:"totalPages"`
			PagingCounter int    `json:"pagingCounter"`
			PrevPage      string `json:"prevPage"`
			NextPage      string `json:"nextPage"`
		} `json:"pagination"`
		Balance []struct {
			Periodic struct {
				AvailableBalance []interface{} `json:"available_balance"`
				LedgerBalance    []interface{} `json:"ledger_balance"`
			} `json:"periodic"`
			Connected []interface{} `json:"connected"`
			Owner     []string      `json:"owner"`
			Record    []string      `json:"record"`
			ID        string        `json:"_id"`
			Account   struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"account"`
			V                int       `json:"__v"`
			AvailableBalance int       `json:"available_balance"`
			CreatedAt        time.Time `json:"created_at"`
			Currency         string    `json:"currency"`
			Customer         struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Env           string    `json:"env"`
			LastUpdated   time.Time `json:"last_updated"`
			LedgerBalance float64   `json:"ledger_balance"`
		} `json:"balance"`
	} `json:"data"`
}

/*
Transaction response
{'status': 'success', 'message': 'Transaction Sets successfully retrieved',
	'data': {'pagination': None, 'trans':
		[{'bank':
			{'_id': '5d6fe57a4099cc4b210bbeb1', 'name': 'First Bank of Nigeria'}, 'count': 79, 'env': 'production-sandbox', 'customer':
			{'_id': '5ee07701a708af0c89a5463c', 'name': 'Joseph Johnson'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbeb8', 'name': 'Sterling Bank'},'count': 59, 'env': 'production-sandbox', 'customer':
		{'_id': '5eae1ec8a400e208b3b888cd', 'name': 'Faith Patterson'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbebb', 'name': 'United Bank For Africa'}, 'count': 52, 'env': 'production-sandbox', 'customer':
		{'_id': '5defaec14ff9b30ae366483a', 'name': 'Godwill Savage'}}, {'bank': {'_id': '5d6fe57a4099cc4b210bbec2', 'name': 'Zenith Bank'},
		'count': 30, 'env': 'production-sandbox', 'customer': {'_id': '5e9d5dd3471ff50f735ad68a', 'name': 'Emmanuel Johnson'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbebc', 'name': 'Wema Bank'}, 'count': 66, 'env': 'production-sandbox', 'customer':
		{'_id': '5ec1df66cd5bde37d3238828', 'name': 'Maryam Savage'}}, {'bank': {'_id': '5d6fe57a4099cc4b210bbebb', 'name': 'United Bank For Africa'},
		'count': 81, 'env': 'production-sandbox', 'customer': {'_id': '5defaec14ff9b30ae366483a', 'name': 'Godwill Savage'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbeba', 'name': 'Keystone Bank'}, 'count': 49, 'env': 'production-sandbox', 'customer':
		{'_id': '5ec552e534dce7257892ed79', 'name': 'Christina Patterson'}}, {'bank': {'_id': '5d6fe57a4099cc4b210bbeb3', 'name': 'Guaranty Trust Bank'},
		'count': 30, 'env': 'production-sandbox', 'customer': {'_id': '5ed3d67b8723c11444c43283', 'name': 'Rachel Simpson'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbeb3', 'name': 'Guaranty Trust Bank'}, 'count': 33, 'env': 'production-sandbox', 'customer':
		{'_id': '5ed3d67b8723c11444c43283', 'name': 'Rachel Simpson'}}, {'bank': {'_id': '5d6fe57a4099cc4b210bbeb6', 'name': 'Stanbic IBTC Bank'},
		'count': 19, 'env': 'production-sandbox', 'customer': {'_id': '5df75c17c4e02270a383fb0e', 'name': 'Sunday Pitt'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbeb8', 'name': 'Sterling Bank'}, 'count': 56, 'env': 'production-sandbox', 'customer':
		{'_id': '5eae1ec8a400e208b3b888cd', 'name': 'Faith Patterson'}}, {'bank': {'_id': '5d6fe57a4099cc4b210bbec0', 'name': 'Access Bank'},
		'count': 65, 'env': 'production-sandbox', 'customer': {'_id': '5ede59514eb21a091d146ebf', 'name': 'Endurance Agusto'}},
		{'bank': {'_id': '5d6fe57a4099cc4b210bbeb3', 'name': 'Guaranty Trust Bank'}, 'count': 69, 'env': 'production-sandbox',
		'customer': {'_id': '5ed3d67b8723c11444c43283', 'name': 'Rachel Simpson'}}]}
}
*/
