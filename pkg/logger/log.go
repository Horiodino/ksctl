package logger

import (
	"fmt"
	"io"
	"log/slog"
	"math"
	"strings"

	box "github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
	"github.com/ksctl/ksctl/pkg/helpers/utilities"
	"github.com/ksctl/ksctl/pkg/resources"
	cloudController "github.com/ksctl/ksctl/pkg/resources/controllers/cloud"
	"github.com/rodaine/table"
)

type Logger struct {
	logger     *slog.Logger
	moduleName string
}

const (
	LimitCol = 80
)

func (l *Logger) SetPackageName(m string) {
	l.moduleName = m
}

func newLogger(out io.Writer, ver slog.Level, debug bool) *slog.Logger {
	if debug {
		return slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: ver,
		}))
	}
	return slog.New(slog.NewTextHandler(out, &slog.HandlerOptions{
		Level: ver,
	}))
}

func NewDefaultLogger(verbose int, out io.Writer) resources.LoggerFactory {
	// LevelDebug Level = -4
	// LevelInfo  Level = 0
	// LevelWarn  Level = 4
	// LevelError Level = 8

	var ve slog.Level

	if verbose < 0 {
		ve = slog.LevelDebug

		return &Logger{logger: newLogger(out, ve, true)}

	} else if verbose < 4 {
		ve = slog.LevelInfo
	} else if verbose < 8 {
		ve = slog.LevelWarn
	} else {
		ve = slog.LevelError
	}

	return &Logger{logger: newLogger(out, ve, false)}
}

func (l *Logger) Print(msg string, args ...any) {
	args = append([]any{"package", l.moduleName}, args...)
	l.logger.Info(msg, args...)
}

func (l *Logger) Success(msg string, args ...any) {
	color.Set(color.FgGreen, color.Bold)
	defer color.Unset()
	args = append([]any{"package", l.moduleName, "msgType", "SUCCESS"}, args...)
	l.logger.Info(msg, args...)
}

func (l *Logger) Note(msg string, args ...any) {
	color.Set(color.FgBlue, color.Bold)
	defer color.Unset()
	args = append([]any{"package", l.moduleName, "msgType", "NOTE"}, args...)
	l.logger.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	defer color.Unset()
	args = append([]any{"package", l.moduleName}, args...)
	l.logger.Debug(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	color.Set(color.FgHiRed, color.Bold)
	defer color.Unset()
	args = append([]any{"package", l.moduleName}, args...)

	msgStrErr := fmt.Sprintf("%v", args)
	l.logger.Error(msg, "Reason", msgStrErr)
}

func (l *Logger) NewError(format string, args ...any) error {
	l.Debug(format, args...)
	args = append([]any{"err_package", l.moduleName}, args...)
	return fmt.Errorf(format, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	color.Set(color.FgYellow, color.Bold)
	defer color.Unset()
	args = append([]any{"package", l.moduleName, "msgType", "WARN"}, args...)
	l.logger.Warn(msg, args...)
}

func (l *Logger) Table(data []cloudController.AllClusterData) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Provider", "Nodes", "Type", "K8s")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, row := range data {
		node := ""
		if row.Type == "ha" {
			node = fmt.Sprintf("cp: %d\nwp: %d\nds: %d\nlb: 1", row.NoCP, row.NoWP, row.NoDS)
		} else {
			node = fmt.Sprintf("wp: %d", row.NoMgt)
		}
		tbl.AddRow(row.Name, string(row.Provider)+"("+row.Region+")", node, row.Type, string(row.K8sDistro))
	}

	tbl.Print()
}

func addLineTerminationForLongStrings(str string) string {

	//arr with endline split
	arrStr := strings.Split(str, "\n")

	var helper func(string) string

	helper = func(_str string) string {

		if len(_str) < LimitCol {
			return _str
		}

		x := string(utilities.DeepCopySlice[byte]([]byte(_str[:LimitCol])))
		y := string(utilities.DeepCopySlice[byte]([]byte(helper(_str[LimitCol:]))))

		// ks
		// ^^
		if x[len(x)-1] != ' ' && y[0] != ' ' {
			x += "-"
		}

		_new := x + "\n" + y
		return _new
	}

	for idx, line := range arrStr {
		arrStr[idx] = helper(line)
	}

	return strings.Join(arrStr, "\n")
}

func (l *Logger) Box(title string, lines string) {
	px := 4

	if len(title) >= 2*px+len(lines) {
		// some maths
		px = int(math.Ceil(float64(len(title)-len(lines))/2)) + 1
	}

	l.Debug("PostUpdate Box", "px", px, "title", len(title), "lines", len(lines))

	Box := box.New(box.Config{
		Px:       px,
		Py:       2,
		Type:     "Bold",
		TitlePos: "Top",
		Color:    "Cyan"})

	Box.Println(title, addLineTerminationForLongStrings(lines))
}
