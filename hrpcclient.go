package hivego

import (
    "encoding/json"
    "errors"
    "github.com/cfoxon/jsonrpc2client"
    "log"
    "strconv"
)

type HiveRpcNode struct {
    address string
    MaxConn int
    MaxBatch int
    NoBroadcast bool
}

type globalProps struct {
    HeadBlockNumber int `json:"head_block_number"`
    HeadBlockId string `json:"head_block_id"`
    Time string `json:"time"`
}

type hrpcQuery struct {
    method string
    params interface{}
}

func NewHiveRpc(addr string) *HiveRpcNode {
    return NewHiveRpcWithOpts(addr, 1, 4)
}

func NewHiveRpcWithOpts(addr string, maxConn int, maxBatch int) *HiveRpcNode {
    return &HiveRpcNode{address: addr,
        MaxConn: maxConn,
        MaxBatch: maxBatch,
    }
}

func (h *HiveRpcNode) GetDynamicGlobalProps() ([]byte, error){
    q := hrpcQuery{method: "condenser_api.get_dynamic_global_properties", params: []string{}}
    res, err := h.rpcExec(h.address, q)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (h *HiveRpcNode) rpcExec(endpoint string, query hrpcQuery) ([]byte, error) {
    rpcClient := jsonrpc2client.NewClientWithOpts(endpoint, h.MaxConn, h.MaxBatch)
    jr2query := &jsonrpc2client.RpcRequest{Method: query.method,  JsonRpc: "2.0", Id: 1, Params: query.params}
    resp, err :=rpcClient.CallRaw(jr2query)
    if err != nil {
        return nil, err
    }

    if resp.Error != nil {
        return nil, errors.New(strconv.Itoa(resp.Error.Code) + "    " + resp.Error.Message)
    }

    return resp.Result, nil
}

func (h *HiveRpcNode) rpcExecBatch(endpoint string, queries []hrpcQuery) ([]json.RawMessage, error) {
    rpcClient := jsonrpc2client.NewClientWithOpts(endpoint, h.MaxConn, h.MaxBatch)

    var jr2queries jsonrpc2client.RPCRequests
    for i, query := range queries {
        jr2query := &jsonrpc2client.RpcRequest{Method: query.method, JsonRpc: "2.0", Id: i, Params: query.params}
        jr2queries = append(jr2queries, jr2query)
    }

    resps, err :=rpcClient.CallBatchRaw(jr2queries)
    if err != nil {
        return nil, err
    }

    var batchResult []json.RawMessage
    for _, resp := range resps {
        thisResult := json.RawMessage{}
        if err := json.Unmarshal(resp.Result, &thisResult); err != nil {
            log.Println("err unmarshalling res.result")
            log.Println(err)
            log.Println(resp)
        }
        batchResult = append(batchResult, thisResult)
    }

    return batchResult, nil
}

func (h *HiveRpcNode) rpcExecBatchFast(endpoint string, queries []hrpcQuery) ([][]byte, error) {
    rpcClient := jsonrpc2client.NewClientWithOpts(endpoint, h.MaxConn, h.MaxBatch)

    var jr2queries jsonrpc2client.RPCRequests
    for i, query := range queries {
        jr2query := &jsonrpc2client.RpcRequest{Method: query.method, JsonRpc: "2.0", Id: i, Params: query.params}
        jr2queries = append(jr2queries, jr2query)
    }

    resps, err :=rpcClient.CallBatchFast(jr2queries)
    if err != nil {
        return nil, err
    }

    var batchResult [][]byte
    for _, resp := range resps {
        batchResult = append(batchResult, resp)
    }

    return batchResult, nil
}