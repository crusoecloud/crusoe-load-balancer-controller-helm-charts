package internal

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"gitlab.com/crusoeenergy/schemas/utils/log"
)

// TODO(template) update or delete this. See README.md.
func RunCmd() {
	// Listen for interrupt signals.
	interrupt := make(chan os.Signal, 1)
	// Ctrl-C
	signal.Notify(interrupt, os.Interrupt)
	// this is what docker sends when shutting down a container
	signal.Notify(interrupt, syscall.SIGTERM)

	// TODO(template) pass ctx to any long running goroutines started here so they can shutdown gracefully.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			return

		case <-interrupt:
			// TODO(template) set service name
			log.Info(ctx).Msg("interrupt signal received - shutting down <service>")
			cancel()
		}
	}()
	// TODO(template) this will force shutdown hangDetectionSleepTime after context is canceled. Remove if unwanted
	go detectAndKillHang(ctx)

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z07:00" // RFC3339 with millisecond precision

	log.Info(ctx).Msg("Hello, world!")
}

const (
	hangDetectionSleepTime = 30 * time.Second
)

func detectAndKillHang(ctx context.Context) {
	<-ctx.Done()
	time.Sleep(hangDetectionSleepTime)

	// Print all goroutine stack traces
	buf := make([]byte, 1<<16) //nolint:gomnd // pick a big enough buffer size to capture all goroutines
	bytesWritten := runtime.Stack(buf, true)
	log.Info(ctx).Msgf("%s", buf[:bytesWritten])

	panic("hang detected")
}
