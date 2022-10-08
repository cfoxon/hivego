package hivego

type TransactionQueryParams struct {
	TransactionId     string `json:"id"`
	IncludeReversible bool   `json:"include_reversible"`
}

func (h *HiveRpcNode) GetTransaction(txId string, includeReversible bool) ([]byte, error) {
	var query = hrpcQuery{method: "account_history_api.get_transaction", params: TransactionQueryParams{TransactionId: txId, IncludeReversible: includeReversible}}
	endpoint := h.address
	res, err := h.rpcExec(endpoint, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
