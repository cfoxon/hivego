# HiveGo - A client for the Hive blockchain

At this time, there are only a few functions from the client. More will be added.

### Example usage:
create a client:
```
hrpc := hivego.NewHiveRpc("https://api.myHiveBlockcchainNode.com")
```

submit a custom json tx:
```
txid, err := hrpc.BroadcastJson([]string{submittingAccount}, []string{}, id, string(jsonPayload), &activeWif)
```

vote a post:
```
txid, err := hrpc.VotePost(voter, author, permlink, weight, &wif)
```

get n blocks starting from block x as the raw response from the rpc (in bytes):
```
responseBytes, err := hrpc.GetBlockRangeFast(startBlock int, count int)
```
WARNING: It is not recommended to stream blocks from public APIs. They are provided as a service to users and saturating them with block requests may (rightfully) result in your IP getting banned
