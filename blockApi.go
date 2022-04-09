package hivego

import (
    "encoding/json"
)

type getBlockRangeQueryParams struct {
    StartingBlockNum int `json:"starting_block_num"`
    Count            int `json:"count"`

}

func (h HiveRpcNode) GetBlockRange(startBlock int, count int) ([]json.RawMessage, error) {
    if h.MaxConn == 0 {
        h.MaxConn = 10
    }
    if h.MaxBatch == 0 {
        h.MaxBatch = 4
    }

    var queries []hrpcQuery
    for i := startBlock; i <= startBlock+count; {
        params := getBlockRangeQueryParams{StartingBlockNum: i, Count: 500}
        query := hrpcQuery{method: "block_api.get_block_range", params: params}
        queries = append(queries, query)
        i += 500
    }

    endpoint := h.address

    res, err := h.rpcExecBatch(endpoint,queries)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func (h HiveRpcNode) GetBlockRangeFast(startBlock int, count int) ([][]byte, error) {
    if h.MaxConn == 0 {
        h.MaxConn = 10
    }
    if h.MaxBatch == 0 {
        h.MaxBatch = 4
    }

    var queries []hrpcQuery
    for i := startBlock; i <= startBlock+count; {
        params := getBlockRangeQueryParams{StartingBlockNum: i, Count: 500}
        query := hrpcQuery{method: "block_api.get_block_range", params: params}
        queries = append(queries, query)
        i += 500
    }

    endpoint := h.address

    res, err := h.rpcExecBatchFast(endpoint,queries)
    if err != nil {
        return nil, err
    }

    return res, nil
}
