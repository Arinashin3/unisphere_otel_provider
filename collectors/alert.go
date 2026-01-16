package collectors

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"time"
	"unisphere_otel_provider/gounity/api"
	"unisphere_otel_provider/utils/enum"

	"github.com/tidwall/gjson"
	"go.opentelemetry.io/otel/log"
)

func init() {
	key := "alert"
	registerModule(key, NewAlert())

}

type ModuleAlert struct {
	name      string
	opts      *api.UnityActionOptions
	defaults  bool
	timestamp time.Time

	// Configuration File
	Enabled *bool `yaml:"enabled,omitempty"`
	Level   int64 `yaml:"level,omitempty"`
}

func NewAlert() *ModuleAlert {
	return &ModuleAlert{
		defaults: false,
		Level:    0,
	}
}

func (_m *ModuleAlert) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

func (_m *ModuleAlert) Init(key string) {
	_m.name = key
	_m.opts = api.NewUnityActionOptions("alert")
	_m.opts.Fields = []string{
		"timestamp",
		"severity",
		"messageId",
		"message",
	}
}

func (_m *ModuleAlert) Run(logger *slog.Logger, col *Collector) {
	opt := *_m.opts
	ctime := time.Now().Add(-1 * time.Hour).UTC()
	client := col.Client
	lp := col.LoggerProvider

	for {
		pvlogger := lp.Logger(_m.name, log.WithInstrumentationAttributes(col.detectLabels...))
		opt.Filters = []string{
			"timestamp gt \"" + ctime.Format("2006-01-02T15:04:05.000Z") + "\"",
		}

		tmpTime := time.Now().UTC()
		data, err := client.GetInstances(&opt)
		if err != nil {
			logger.Error("Error to GET AlertLog", "err", err)
			col.success = false
			time.Sleep(col.interval)
			continue
		}
		col.success = true
		if len(data) == 0 {
			time.Sleep(col.interval)
			continue
		}

		for _, v := range data {
			record := log.Record{}
			if _m.Level > v.Get("severity").Int() {
				continue
			}

			record.SetTimestamp(v.Get("creationTime").Time())
			logBody := struct {
				Message   string `json:"message"`
				MessageId string `json:"message_id"`
			}{
				v.Get("message").String(),
				v.Get("messageId").String(),
			}
			record.SetTimestamp(v.Get("timestamp").Time())
			jsonBody, _ := json.Marshal(logBody)
			body := gjson.ParseBytes(jsonBody).String()
			record.SetBody(log.StringValue(body))
			record.AddAttributes(
				log.String("level", enum.SeverityEnum(v.Get("severity").Int()).String()),
			)
			pvlogger.Emit(col.ctx, record)

		}
		ctime = tmpTime
		time.Sleep(col.interval)
	}

}
