package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"thermal-throttling-analyzer/internal/analyzer"
	"thermal-throttling-analyzer/internal/events"
	"thermal-throttling-analyzer/internal/sensors"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Tell me when things go bad",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring thermal state... (Press Ctrl+C to stop)")

		sm := analyzer.NewStateMachine()
		logger, err := events.NewLogger()
		if err != nil {
			fmt.Printf("Error initializing logger: %v\n", err)
			return
		}

		// Channel for clean exit
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		ticker := time.NewTicker(analyzer.SampleInterval) // 2s
		defer ticker.Stop()

		lastState := analyzer.StateNormal
		
		for {
			select {
			case <-sigChan:
				fmt.Println("\nStopping monitor.")
				return
			case <-ticker.C:
				snap := sensors.CollectSnapshot()
				res := sm.UpdateWithHistory(snap)

				// Print only on state change or significant event?
				// Requirement: "Prints output only on state change"
				
				if res.State != lastState {
					timestamp := time.Now().Format("15:04")
					fmt.Printf("[%s] %s detected (temp %.0f°C)\n", timestamp, res.State, snap.TempC)
					
					// Log event
					_ = logger.LogEvent(events.Event{
						Timestamp: time.Now(),
						Type:      string(res.State), // Log the state change
						State:     string(res.State),
						Details:   res.Reason,
					})
					
					lastState = res.State
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
