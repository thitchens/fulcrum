package log

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func Error(message string) {
	Logger.Error(message)
}

func Fatal(message string) {
	Logger.Fatal(message)
}

func Init() error {
	var err error
	Logger, err = buildLogger()
	return err
}

func buildLogger() (*zap.Logger, error) {
	// TODO figure out how to trim build path from stack trace
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"fulcrum.log", "stderr"}
	return cfg.Build()
}
