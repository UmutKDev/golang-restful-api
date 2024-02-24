package Config

import (
	Nested "github.com/antonfisher/nested-logrus-formatter"
	Log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() {

	Log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	Log.SetReportCaller(true)
	Log.SetFormatter(&Nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})

}

func getLoggerLevel(value string) Log.Level {
	switch value {
	case "DEBUG":
		return Log.DebugLevel
	case "TRACE":
		return Log.TraceLevel
	default:
		return Log.InfoLevel
	}
}
