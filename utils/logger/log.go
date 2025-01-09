package logger

import (
	"sync"
	"time"

	ierrors "movie-app/utils/errors"

	"movie-app/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var singleton sync.Once

func InitLogger(cfg config.Config) {
	singleton.Do(func() {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)

		zerolog.ErrorStackMarshaler = ierrors.MarshalStack
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.RFC3339
		})

		logger := log.With().Stack().Logger().Output(output)
		log.Logger = logger
	})
}
