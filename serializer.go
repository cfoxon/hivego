package hivego

import (
    "bytes"
    "encoding/binary"
    "time"
)
func opIdB(opName string) byte {
    id := getHiveOpId(opName)
    return byte(id)
}

func refBlockNumB(refBlockNumber uint16) []byte {
    buf := make([]byte, 2)
    binary.LittleEndian.PutUint16(buf, refBlockNumber)
    return buf
}

func refBlockPrefixB(refBlockPrefix uint32) []byte {
    buf := make([]byte, 4)
    binary.LittleEndian.PutUint32(buf, refBlockPrefix)
    return buf
}

func expTimeB(expTime string) ([]byte, error) {
    exp, err := time.Parse("2006-01-02T15:04:05", expTime)
    if err != nil {
        return nil, err
    }
    buf := make([]byte, 4)
    binary.LittleEndian.PutUint32(buf, uint32(exp.Unix()))
    return buf, nil
}

func countOpsB(ops []hiveOperation) []byte {
    b := make([]byte, 5)
    l := binary.PutUvarint(b, uint64(len(ops)))
    return b[0:l]
}

func extensionsB() byte {
    return byte(0x00)
}

func appendVString(s string, b *bytes.Buffer) *bytes.Buffer {
    vBuf := make([]byte, 5)
    vLen := binary.PutUvarint(vBuf, uint64(len(s)))
    b.Write(vBuf[0:vLen])

    b.WriteString(s)
    return b
}

func appendVStringArray(a []string, b *bytes.Buffer) *bytes.Buffer {
    b.Write([]byte{byte(len(a))})
    for _, s := range a {
        appendVString(s, b)
    }
    return b
}

func serializeTx(tx hiveTransaction) ([]byte, error) {
    var buf bytes.Buffer
    buf.Write(refBlockNumB(tx.RefBlockNum))
    buf.Write(refBlockPrefixB(tx.RefBlockPrefix))
    expTime, err := expTimeB(tx.Expiration)
    if err != nil {
        return nil, err
    }
    buf.Write(expTime)

    opsB, err := serializeOps(tx.Operations)
    if err != nil {
        return nil, err
    }
    buf.Write(opsB)
    buf.Write([]byte{extensionsB()})
    return buf.Bytes(), nil
}

func serializeOps(ops []hiveOperation) ([]byte, error) {
    var opsBuf bytes.Buffer
    opsBuf.Write(countOpsB(ops))
    for _, op := range ops {
        b, err := op.serializeOp()
        if err != nil {
            return nil, err
        }
        opsBuf.Write(b)
    }
    return opsBuf.Bytes(), nil
}

func (o voteOperation) serializeOp() ([]byte, error){
    var voteBuf bytes.Buffer
    voteBuf.Write([]byte{opIdB(o.opText)})
    appendVString(o.Voter, &voteBuf)
    appendVString(o.Author, &voteBuf)
    appendVString(o.Permlink, &voteBuf)

    weightBuf := make([]byte, 2)
    binary.LittleEndian.PutUint16(weightBuf, uint16(o.Weight))
    voteBuf.Write(weightBuf)

    return voteBuf.Bytes(), nil
}

func (o customJsonOperation) serializeOp() ([]byte, error) {
    var jBuf bytes.Buffer
    jBuf.Write([]byte{opIdB(o.opText)})
    appendVStringArray(o.RequiredAuths, &jBuf)
    appendVStringArray(o.RequiredPostingAuths, &jBuf)
    appendVString(o.Id, &jBuf)
    appendVString(o.Json, &jBuf)

    return jBuf.Bytes(), nil
}

