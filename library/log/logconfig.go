package log

type LogConfig struct {
	LogLevel       string
	UseLogRotation bool
	LogProps       LogProps
}
type LogProps struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}
