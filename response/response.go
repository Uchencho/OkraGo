package response

import "time"

type RetrieveAuthPayload struct {
	StatusCode int    `json:"statusCode"`
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
	StatusCode int    `json:"statusCode"`
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

type AuthByCustomerIDPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
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

type AuthByDateRangePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
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

type AuthByCustomerDateRangePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
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

type AuthByBankPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
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

// --------------
// --------------

type RetrieveBalancePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination interface{} `json:"pagination"`
		Balances   []struct {
			ID      string `json:"id"`
			Account struct {
				ID   string `json:"_id"`
				Bank struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"bank"`
				Nuban string `json:"nuban"`
				Name  string `json:"name"`
			} `json:"account"`
			LedgerBalance    float64   `json:"ledger_balance"`
			AvailableBalance int       `json:"available_balance"`
			Currency         string    `json:"currency"`
			Connected        bool      `json:"connected"`
			Env              string    `json:"env"`
			CreatedAt        time.Time `json:"created_at"`
			LastUpdated      time.Time `json:"last_updated"`
		} `json:"balances"`
	} `json:"data"`
}

type BalanceByIDPayload struct {
	StatusCode int    `json:"statusCode"`
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

type BalanceByCustomerIDPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Balance []struct {
			Periodic struct {
				AvailableBalance []interface{} `json:"available_balance"`
				LedgerBalance    []interface{} `json:"ledger_balance"`
			} `json:"periodic"`
			Connected []string `json:"connected"`
			Owner     []string `json:"owner"`
			Record    []string `json:"record"`
			ID        string   `json:"_id"`
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

type BalanceByCustomerDateRangePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Balance []struct {
			Periodic struct {
				AvailableBalance []interface{} `json:"available_balance"`
				LedgerBalance    []interface{} `json:"ledger_balance"`
			} `json:"periodic"`
			Connected []string `json:"connected"`
			Record    []string `json:"record"`
			ID        string   `json:"_id"`
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

type BalanceByAccountPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Balance []struct {
			Periodic struct {
				AvailableBalance []interface{} `json:"available_balance"`
				LedgerBalance    []interface{} `json:"ledger_balance"`
			} `json:"periodic"`
			Connected []string `json:"connected"`
			Owner     []string `json:"owner"`
			Record    []string `json:"record"`
			ID        string   `json:"_id"`
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

type PeriodicBalancePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		CallbackURL string `json:"callback_url"`
	} `json:"data"`
}

// ----------
// ----------

type RetrieveTransactionPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination interface{} `json:"pagination"`
		Trans      []struct {
			Bank struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"bank"`
			Count    int    `json:"count"`
			Env      string `json:"env"`
			Customer struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
		} `json:"trans"`
	} `json:"data"`
}

type TransactionByCustomerIDPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Transaction []struct {
			Notes struct {
				Topics        []string `json:"topics"`
				Places        []string `json:"places"`
				People        []string `json:"people"`
				Actions       []string `json:"actions"`
				Subjects      []string `json:"subjects"`
				Preopositions []string `json:"preopositions"`
				Desc          string   `json:"desc"`
			} `json:"notes"`
			Fetched                []string      `json:"fetched"`
			Record                 []string      `json:"record"`
			Analyzed               []interface{} `json:"analyzed"`
			Ner                    interface{}   `json:"ner"`
			NerV                   int           `json:"ner_v"`
			ID                     string        `json:"_id"`
			Branch                 string        `json:"branch"`
			ClearedDate            time.Time     `json:"cleared_date"`
			UnformattedClearedDate string        `json:"unformatted_cleared_date"`
			Code                   string        `json:"code"`
			Credit                 int           `json:"credit"`
			Debit                  interface{}   `json:"debit"`
			TransDate              time.Time     `json:"trans_date"`
			UnformattedTransDate   string        `json:"unformatted_trans_date"`
			Customer               struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Account struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"account"`
			Env  string `json:"env"`
			Bank struct {
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Icon   string `json:"icon"`
				Status string `json:"status"`
			} `json:"bank"`
			CreatedAt   time.Time `json:"created_at"`
			LastUpdated time.Time `json:"last_updated"`
			V           int       `json:"__v"`
		} `json:"transaction"`
	} `json:"data"`
}

type TransactionByDateRangePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      int         `json:"nextPage"`
		} `json:"pagination"`
		Transaction []struct {
			Notes struct {
				Topics        []string `json:"topics"`
				Places        []string `json:"places"`
				People        []string `json:"people"`
				Actions       []string `json:"actions"`
				Subjects      []string `json:"subjects"`
				Preopositions []string `json:"preopositions"`
				Desc          string   `json:"desc"`
			} `json:"notes"`
			Fetched                []string      `json:"fetched"`
			Record                 []string      `json:"record"`
			Analyzed               []interface{} `json:"analyzed"`
			Ner                    interface{}   `json:"ner"`
			NerV                   int           `json:"ner_v"`
			ID                     string        `json:"_id"`
			Branch                 string        `json:"branch"`
			ClearedDate            time.Time     `json:"cleared_date"`
			UnformattedClearedDate string        `json:"unformatted_cleared_date"`
			Code                   string        `json:"code"`
			Credit                 interface{}   `json:"credit"`
			Debit                  int           `json:"debit"`
			TransDate              time.Time     `json:"trans_date"`
			UnformattedTransDate   string        `json:"unformatted_trans_date"`
			Customer               struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Account struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"account"`
			Env  string `json:"env"`
			Bank struct {
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Icon   string `json:"icon"`
				Status string `json:"status"`
			} `json:"bank"`
			CreatedAt   time.Time `json:"created_at"`
			LastUpdated time.Time `json:"last_updated"`
			V           int       `json:"__v"`
		} `json:"transaction"`
	} `json:"data"`
}

type TransactionByCustomerDateRangePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      int         `json:"nextPage"`
		} `json:"pagination"`
		Transaction []struct {
			Notes struct {
				Topics        []string `json:"topics"`
				Places        []string `json:"places"`
				People        []string `json:"people"`
				Actions       []string `json:"actions"`
				Subjects      []string `json:"subjects"`
				Preopositions []string `json:"preopositions"`
				Desc          string   `json:"desc"`
			} `json:"notes"`
			Fetched                []string      `json:"fetched"`
			Record                 []string      `json:"record"`
			Analyzed               []interface{} `json:"analyzed"`
			Ner                    interface{}   `json:"ner"`
			NerV                   int           `json:"ner_v"`
			ID                     string        `json:"_id"`
			Branch                 string        `json:"branch"`
			ClearedDate            time.Time     `json:"cleared_date"`
			UnformattedClearedDate string        `json:"unformatted_cleared_date"`
			Code                   string        `json:"code"`
			Credit                 int           `json:"credit"`
			Debit                  interface{}   `json:"debit"`
			TransDate              time.Time     `json:"trans_date"`
			UnformattedTransDate   string        `json:"unformatted_trans_date"`
			Customer               struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Account struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"account"`
			Env  string `json:"env"`
			Bank struct {
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Icon   string `json:"icon"`
				Status string `json:"status"`
			} `json:"bank"`
			CreatedAt   time.Time `json:"created_at"`
			LastUpdated time.Time `json:"last_updated"`
			V           int       `json:"__v"`
		} `json:"transaction"`
	} `json:"data"`
}

// -------------------
// -------------------

type IdentityByCustomerDatePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Identity []interface{} `json:"identity"`
	} `json:"data"`
}

// -------------------
// -------------------

type IncomeByCustomerDatePayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Pagination struct {
			TotalDocs     int         `json:"totalDocs"`
			Limit         int         `json:"limit"`
			HasPrevPage   bool        `json:"hasPrevPage"`
			HasNextPage   bool        `json:"hasNextPage"`
			Page          int         `json:"page"`
			TotalPages    int         `json:"totalPages"`
			PagingCounter int         `json:"pagingCounter"`
			PrevPage      interface{} `json:"prevPage"`
			NextPage      interface{} `json:"nextPage"`
		} `json:"pagination"`
		Income []interface{} `json:"income"`
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

/*
Retrieve auth response
{"status":"success","message":"Auths successfully retrieved","data":{"pagination":null,"auths":[{"id":"5f231e0c9e5c6e823adf601d","validated":true,"record":{"_id":"5f231df19f28af16b6cf5994","bank":{"_id":"5d6fe57a4099cc4b210bbeae","name":"Ecobank Nigeria"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5df1020beffe401cfaaa473f","name":"James Galler"}},"created_at":"2020-07-30T19:22:52.076Z","last_updated":"2020-07-30T19:22:52.076Z"},{"id":"5f03b2d4af4681496afdd7b4","validated":true,"record":{"_id":"5f03b2c28b1d840dc57b7f41","bank":{"_id":"5d6fe57a4099cc4b210bbeb9","name":"Union Bank of Nigeria"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ef09108658acd0d06ae06b9","name":"Jorge Agusto"}},"created_at":"2020-07-06T23:25:08.750Z","last_updated":"2020-07-06T23:25:08.750Z"},{"id":"5f03aea1af4681496afc3c88","validated":true,"record":{"_id":"5f03ae86ac55690d56c48c92","bank":{"_id":"5d6fe57a4099cc4b210bbebc","name":"Wema Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ec1df66cd5bde37d3238828","name":"Maryam Savage"}},"created_at":"2020-07-06T23:07:13.092Z","last_updated":"2020-07-06T23:07:13.092Z"},{"id":"5f03ae0baf4681496afbcb42","validated":true,"record":{"_id":"5f03adf6414c210d08ed6f2d","bank":{"_id":"5d6fe57a4099cc4b210bbebb","name":"United Bank For Africa"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5defaec14ff9b30ae366483a","name":"Godwill Savage"}},"created_at":"2020-07-06T23:04:43.370Z","last_updated":"2020-07-06T23:04:43.370Z"},{"id":"5f03aca9af4681496afb4f4c","validated":true,"record":{"_id":"5f03ac9ce167d80cd2514368","bank":{"_id":"5d6fe57a4099cc4b210bbec2","name":"Zenith Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5e9d5dd3471ff50f735ad68a","name":"Emmanuel Johnson"}},"created_at":"2020-07-06T22:58:49.621Z","last_updated":"2020-07-06T22:58:49.621Z"},{"id":"5f03a8cdaf4681496afaccb9","validated":true,"record":{"_id":"5f03a8b95a63fd0c4a54cacb","bank":{"_id":"5d6fe57a4099cc4b210bbeb8","name":"Sterling Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5eae1ec8a400e208b3b888cd","name":"Mary Galler"}},"created_at":"2020-07-06T22:42:21.904Z","last_updated":"2020-07-06T22:42:21.904Z"},{"id":"5f039c38af4681496af81d30","validated":true,"record":{"_id":"5f039c243103140b668e92fb","bank":{"_id":"5d6fe57a4099cc4b210bbeb3","name":"Guaranty Trust Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ed3d67b8723c11444c43283","name":"Rachel Simpson"}},"created_at":"2020-07-06T21:48:40.407Z","last_updated":"2020-07-06T21:48:40.407Z"},{"id":"5f004eaeaf4681496a14aad4","validated":true,"record":{"_id":"5f004e4eef5d3058b46463a8","bank":{"_id":"5d6fe57a4099cc4b210bbeba","name":"Keystone Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ec552e534dce7257892ed79","name":"Christina Patterson"}},"created_at":"2020-07-04T09:41:02.528Z","last_updated":"2020-07-04T09:41:02.528Z"},{"id":"5eede9aeb87f052d8dc5163e","validated":true,"record":{"_id":"5eede99e5156602b05620833","bank":{"_id":"5d6fe57a4099cc4b210bbeb3","name":"Guaranty Trust Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ed3d67b8723c11444c43283","name":"Rachel Simpson"}},"created_at":"2020-06-20T10:49:18.524Z","last_updated":"2020-06-20T10:49:18.524Z"},{"id":"5ee7cc255a297caf4c6d5704","validated":true,"record":{"_id":"5ee7cc1678b6e314502cdca2","bank":{"_id":"5d6fe57a4099cc4b210bbec0","name":"Access Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ede59514eb21a091d146ebf","name":"Endurance Agusto"}},"created_at":"2020-06-15T19:29:41.739Z","last_updated":"2020-06-15T19:29:41.739Z"},{"id":"5ee63c5d5a297caf4c33302d","validated":true,"record":{"_id":"5ee63c4f98584109765ba885","bank":{"_id":"5d6fe57a4099cc4b210bbeb6","name":"Stanbic IBTC Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5df75c17c4e02270a383fb0e","name":"Sunday Pitt"}},"created_at":"2020-06-14T15:03:57.629Z","last_updated":"2020-06-14T15:03:57.629Z"},{"id":"5ee5029e5a297caf4ce25eb0","validated":true,"record":{"_id":"5ee5028620ec110c6194f85d","bank":{"_id":"5d6fe57a4099cc4b210bbeb3","name":"Guaranty Trust Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ed3d67b8723c11444c43283","name":"Rachel Simpson"}},"created_at":"2020-06-13T16:45:18.846Z","last_updated":"2020-06-13T16:45:18.846Z"},{"id":"5ee4b9195a297caf4cd554ee","validated":true,"record":{"_id":"5ee4b8e9a7226f24e36478c0","bank":{"_id":"5d6fe57a4099cc4b210bbebb","name":"United Bank For Africa"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5defaec14ff9b30ae366483a","name":"Godwill Savage"}},"created_at":"2020-06-13T11:31:37.869Z","last_updated":"2020-06-13T11:31:37.869Z"},{"id":"5ee4b4135a297caf4cd3a011","validated":true,"record":{"_id":"5ee4b3b33d78a02417326ffa","bank":{"_id":"5d6fe57a4099cc4b210bbeb8","name":"Sterling Bank"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5eae1ec8a400e208b3b888cd","name":"Mary Galler"}},"created_at":"2020-06-13T11:10:11.823Z","last_updated":"2020-06-13T11:10:11.823Z"},{"id":"5ee4af985a297caf4ccdfd96","validated":true,"record":{"_id":"5ee4af7496277920f91f0584","bank":{"_id":"5d6fe57a4099cc4b210bbeb1","name":"First Bank of Nigeria"},"env":"production-sandbox","owner":"5eda1df311475d0826ac1850","customer":{"_id":"5ee07701a708af0c89a5463c","name":"Joseph Johnson"}},"created_at":"2020-06-13T10:51:04.066Z","last_updated":"2020-06-13T10:51:04.066Z"}]}}
*/
