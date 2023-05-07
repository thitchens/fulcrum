package log

import "go.uber.org/zap"

var (
	ErrorLogger *zap.Logger
	StdOutLogger *zap.Logger
)

func Error(message string) {
	ErrorLogger.Error(message)
}

func Fatal(message string) {
	ErrorLogger.Fatal(message)
}

func Info(message string) {
	StdOutLogger.Info(message)
}

func Init() error {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"fulcrum.log", "stderr"}
	ErrorLogger, err = cfg.Build()
	if err != nil {
		return err
	}

	cfg.OutputPaths = []string{"fulcrim.log", "stdout"}
	StdOutLogger, err = cfg.Build()
	return err
}
