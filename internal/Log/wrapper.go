package Log

import "fmt"

func Debug(message string, args ...any) {
	message = fmt.Sprintf(message, args...)
	logger.Debug().Msg(message)
}
func Info(message string, args ...any) {
	message = fmt.Sprintf(message, args...)
	logger.Info().Msg(message)
}
func Warn(message string, args ...any) {
	message = fmt.Sprintf(message, args...)
	logger.Warn().Msg(message)
}
func Error(err error, customMessage string, args ...any) error {
	if customMessage != "" {
		customMessage = fmt.Sprintf(customMessage, args...)
	}
	logger.Error().Err(err).Msg(customMessage)
	return fmt.Errorf("%s: %w", customMessage, err)
}
