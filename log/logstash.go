package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

// DefaultTimestampFormat 默认日期解析格式
const DefaultTimestampFormat = `2006-01-02 15:04:05`

// WechatFormatter 日志格式化
type WechatFormatter struct{}

// Format 格式化
func (f *WechatFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["datetime"] = entry.Time.Format(DefaultTimestampFormat)
	entry.Data["channel"] = "golang-wechat"

	// set message field
	v, ok := entry.Data["message"]
	if ok {
		entry.Data["fields.message"] = v
	}
	entry.Data["message"] = entry.Message

	// set level field
	v, ok = entry.Data["level"]
	if ok {
		entry.Data["fields.level"] = v
	}
	entry.Data["level"] = entry.Level.String()

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}
