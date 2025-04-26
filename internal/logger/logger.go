package logger

import (
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

var once sync.Once

func Get(flags ...bool) zerolog.Logger {
	once.Do(func() {
		zerolog.TimestampFieldName = "time"
		zerolog.LevelFieldName = "level"
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
			return file + ":" + strconv.Itoa(line)
		}
		if flags[0] {
			logger = zerolog.New(os.Stdout).
				Level(zerolog.DebugLevel).
				With().
				Timestamp().
				Caller().
				Logger().
				Output(zerolog.ConsoleWriter{Out: os.Stdout})
		} else {
			logger = zerolog.New(os.Stdout).
				Level(zerolog.InfoLevel).
				With().
				Timestamp().
				Caller().
				Logger()
		}
	})

	return logger
}