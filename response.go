package okra

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

type TransactionByBankIDPayload struct {
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

// -------------------
// -------------------

type RetrieveIdentitiesPayload struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       struct {
		Transactions []struct {
			Notes struct {
				Topics        []string      `json:"topics"`
				Places        []interface{} `json:"places"`
				People        []string      `json:"people"`
				Actions       []interface{} `json:"actions"`
				Subjects      []string      `json:"subjects"`
				Preopositions []interface{} `json:"preopositions"`
				Desc          string        `json:"desc"`
			} `json:"notes"`
			Fetched     []string    `json:"fetched"`
			Record      []string    `json:"record"`
			ID          string      `json:"_id"`
			TransDate   time.Time   `json:"trans_date"`
			ClearedDate time.Time   `json:"cleared_date"`
			Debit       interface{} `json:"debit"`
			Credit      int         `json:"credit"`
			Ref         string      `json:"ref"`
			Bank        struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
				Icon string `json:"icon"`
			} `json:"bank"`
			Benefactor interface{} `json:"benefactor"`
			Customer   string      `json:"customer"`
			Account    struct {
				ID    string `json:"_id"`
				Name  string `json:"name"`
				Nuban string `json:"nuban"`
			} `json:"account"`
			Env                    string    `json:"env"`
			CreatedAt              time.Time `json:"created_at"`
			LastUpdated            time.Time `json:"last_updated"`
			V                      int       `json:"__v"`
			UnformattedClearedDate string    `json:"unformatted_cleared_date"`
			UnformattedTransDate   string    `json:"unformatted_trans_date"`
		} `json:"transactions"`
	} `json:"data"`
}

type IdentityByCustomerPayload struct {
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
		Identity []struct {
			Enrollment struct {
				Bank             string `json:"bank"`
				Branch           string `json:"branch"`
				RegistrationDate string `json:"registration_date"`
			} `json:"enrollment"`
			Aliases    []interface{} `json:"aliases"`
			Phone      []string      `json:"phone"`
			Email      []string      `json:"email"`
			Verified   bool          `json:"verified"`
			NextOfKins []interface{} `json:"next_of_kins"`
			Address    []string      `json:"address"`
			Record     interface{}   `json:"record"`
			BvnUpdated bool          `json:"bvn_updated"`
			SmileID    bool          `json:"smileId"`
			RefIds     []interface{} `json:"ref_ids"`
			ID         string        `json:"_id"`
			Bvn        string        `json:"bvn"`
			V          int           `json:"__v"`
			CreatedAt  time.Time     `json:"created_at"`
			Customer   struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"customer"`
			Dob            string      `json:"dob"`
			Env            string      `json:"env"`
			Firstname      string      `json:"firstname"`
			Fullname       string      `json:"fullname"`
			Gender         string      `json:"gender"`
			LastUpdated    time.Time   `json:"last_updated"`
			Lastname       string      `json:"lastname"`
			MaritalStatus  interface{} `json:"marital_status"`
			Middlename     interface{} `json:"middlename"`
			Score          string      `json:"score"`
			LevelOfAccount string      `json:"level_of_account"`
			LgaOfOrigin    string      `json:"lga_of_origin"`
			LgaOfResidence string      `json:"lga_of_residence"`
			Nationality    string      `json:"nationality"`
			PhotoID        []struct {
				ID        string `json:"_id"`
				URL       string `json:"url"`
				ImageType string `json:"image_type"`
				Bank      string `json:"bank"`
			} `json:"photo_id"`
			StateOfOrigin    string `json:"state_of_origin"`
			StateOfResidence string `json:"state_of_residence"`
			WatchListed      string `json:"watch_listed"`
		} `json:"identity"`
	} `json:"data"`
}

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

type ProcessIncomePayload struct {
	StatusCode int  `json:"statusCode"`
	Success    bool `json:"success"`
	Data       struct {
		Income struct {
			CustomerID struct {
				Confidence                          string  `json:"confidence"`
				LastYearIncome                      int     `json:"last_year_income"`
				MaxNumberOfOverlappingIncomeStreams int     `json:"max_number_of_overlapping_income_streams"`
				NumberOfIncomeStreams               int     `json:"number_of_income_streams"`
				ProjectedYearlyIncome               float64 `json:"projected_yearly_income"`
				Streams                             []struct {
					AvgMonthlyIncome float64 `json:"avg_monthly_income"`
					Days             int     `json:"days"`
					Desc             string  `json:"desc,omitempty"`
					Employer         string  `json:"employer"`
					IncomeStartDate  string  `json:"income_start_date"`
					Index            int     `json:"index,omitempty"`
					MonthlyIncome    float64 `json:"monthly_income"`
					Type             string  `json:"type"`
				} `json:"streams"`
			} `json:"5ea2a7f042e8ec1df3fc4291"`
		} `json:"income"`
		Msg string `json:"msg"`
	} `json:"data"`
}
