package hivego

import "time"

func getTestVoteOp() hiveOperation {
    return voteOperation{
        Voter: "xeroc",
        Author: "xeroc",
        Permlink: "piston",
        Weight: 10000,
        opText: "vote",
    }
}

func getTestCustomJsonOp() hiveOperation {
    return customJsonOperation{
        RequiredAuths: []string{},
        RequiredPostingAuths: []string{"xeroc"},
        Id: "test-id",
        Json: "{\"testk\":\"testv\"}",
        opText: "custom_json",
    }
}

func getTwoTestOps() []hiveOperation {
    return []hiveOperation{getTestVoteOp(), getTestCustomJsonOp()}
}

func getTestTx(ops []hiveOperation) hiveTransaction {
    exp, _ := time.Parse("2006-01-02T15:04:05","2016-08-08T12:24:17")
    expStr := exp.Format("2006-01-02T15:04:05")

    return hiveTransaction{
        RefBlockNum: 36029,
        RefBlockPrefix: 1164960351,
        Expiration: expStr,
        Operations: ops,
    }
}

func getTestVoteTx() hiveTransaction {
    return getTestTx([]hiveOperation{getTestVoteOp()})
}

func getTestCustomJsonTx() hiveTransaction {
    return getTestTx([]hiveOperation{getTestCustomJsonOp()})
}

func getTestMultipleOpsTx() hiveTransaction {
    return getTestTx(getTwoTestOps())
}
