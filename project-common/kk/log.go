package kk

import (
	"encoding/json"
	"test.com/project-common/tms"
	"time"
)

type FileMap map[string]any

type KafkaLog struct {
	Type     string
	Action   string
	Time     string
	Msg      string
	Field    FileMap
	FuncName string
}

func Error(err error, funcName string, fieldMap FileMap) []byte {
	kl := KafkaLog{
		Type:     "error",
		Action:   "click",
		Time:     tms.Format(time.Now()),
		Msg:      err.Error(),
		Field:    fieldMap,
		FuncName: funcName,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}

func Info(msg string, funcName string, fieldMap FileMap) []byte {
	kl := KafkaLog{
		Type:     "info",
		Action:   "click",
		Time:     tms.Format(time.Now()),
		Msg:      msg,
		Field:    fieldMap,
		FuncName: funcName,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}
