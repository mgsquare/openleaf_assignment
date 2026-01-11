package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
)

type TextHandler struct {
	writer io.Writer
	level  slog.Level
}

func NewTextHandler(w io.Writer, level slog.Level) *TextHandler {
	return &TextHandler{
		writer: w,
		level:  level,
	}
}

func (h *TextHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= h.level
}

func (h *TextHandler) Handle(_ context.Context, r slog.Record) error {

	timestamp := r.Time.Format("02-01-2006 15:04:05")
	millis := r.Time.Nanosecond() / 1e6

	level := strings.ToUpper(r.Level.String())

	file := "unknown"
	function := ""

	if r.PC != 0 {
		fn := runtime.FuncForPC(r.PC)
		if fn != nil {
			function = filepath.Base(fn.Name())

			filePath, _ := fn.FileLine(r.PC)
			file = filepath.Base(filePath)
		}
	}

	msg := r.Message

	var line string
	if level == "ERROR" && function != "" {
		line = fmt.Sprintf(
			"%s , %03d | %s | %s | %s() | %s\n",
			timestamp, millis, level, msg, function, file,
		)
	} else {
		line = fmt.Sprintf(
			"%s , %03d | %s | %s | %s\n",
			timestamp, millis, level, msg, file,
		)
	}

	_, err := h.writer.Write([]byte(line))
	return err
}

func (h *TextHandler) WithAttrs(_ []slog.Attr) slog.Handler {

	return h
}

func (h *TextHandler) WithGroup(_ string) slog.Handler {
	return h
}
