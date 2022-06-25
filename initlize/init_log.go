package initlize

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLog() error {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	writer := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%s.log", "data"),
		MaxSize:    1000,
		MaxBackups: 5,
		LocalTime:  true,
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Logger().With().Caller().Logger()
	log.Info().Msg("init_log succ")
	return nil
}
