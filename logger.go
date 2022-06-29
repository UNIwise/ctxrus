package ctxrus

import (
	"context"

	"github.com/sirupsen/logrus"
)

type log string

const logKey log = "logger"

var defaultLogger *logrus.Logger

// Globally sets the default logger this is used to set logging settings
// like the level, format, and output.
func SetLogger(log *logrus.Logger) {
	defaultLogger = log
}

// Gets the logger from the context.
// If no logger is added to the context, the default logger is returned.
func GetLogger(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(logKey).(*logrus.Entry)
	if !ok {
		if defaultLogger == nil {
			defaultLogger = logrus.New()
		}
		log = logrus.NewEntry(defaultLogger)
	}

	return log
}

// Adds an entry to the context.
// This is used when the logger is provided from elsewhere.
func AddLogger(ctx context.Context, log *logrus.Entry) context.Context {
	ctx, _ = WithLogger(ctx, log)

	return ctx
}

// Returns the context where fields are added to the logger.
// This is useful for adding fields to the logger but not wanting to log anything.
func AddFields(ctx context.Context, fields logrus.Fields) context.Context {
	log := GetLogger(ctx)
	log = log.WithFields(fields)

	return AddLogger(ctx, log)
}

// Returns the context and logger where the logger is added to the context.
func WithLogger(ctx context.Context, log *logrus.Entry) (context.Context, *logrus.Entry) {
	return context.WithValue(ctx, logKey, log), log
}

// Returns the context and logger where the logger has the fields added to it.
func WithFields(ctx context.Context, fields logrus.Fields) (context.Context, *logrus.Entry) {
	log := GetLogger(ctx)
	log = log.WithFields(fields)

	return WithLogger(ctx, log)
}
