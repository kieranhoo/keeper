package logger

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func LogCallProcedure(method, statusCode string) {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("CALL").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "86",
			Dark:  "86",
		}).
		Foreground(lipgloss.Color("0"))
	// Add a custom style for key `err`
	styles.Keys["method"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["method"] = lipgloss.NewStyle().Bold(true)
	styles.Keys["status_code"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["status_code"] = lipgloss.NewStyle().Bold(true)
	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		// TimeFormat:      time.Kitchen,
		Prefix: "protocol",
	})
	gwlogger.Info("[gRPC]", "method", method, "status_code", statusCode)
}
