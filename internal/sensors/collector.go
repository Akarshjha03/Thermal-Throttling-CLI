package sensors

import (
	"time"
)

// CollectSnapshot gathers data from all sensors and returns a unified Snapshot.
func CollectSnapshot() *Snapshot {
	s := &Snapshot{
		Timestamp:    time.Now(),
		ValidSignals: []string{},
	}

	// 1. Temperature
	temp, err := GetCPUTemperature()
	if err == nil {
		s.TempC = temp
		s.ValidSignals = append(s.ValidSignals, "TempC")
	} else {
		// Mock data for development/demonstration if real sensor fails
		// In a real tool, we might report error or leave as 0.
		// For this strict project, we obey "Missing data must lower confidence, not crash"
		// We leave it as 0/Empty and rely on ValidSignals.
	}

	// 2. Frequency
	freq, err := GetCPUFrequency()
	if err == nil {
		s.FreqMHz = freq.CurrentMHz
		s.BaseFreqMHz = freq.BaseMHz
		s.ValidSignals = append(s.ValidSignals, "FreqMHz", "BaseFreqMHz")
	}

	// 3. Load
	load, err := GetCPULoad()
	if err == nil {
		s.LoadPercent = load
		s.ValidSignals = append(s.ValidSignals, "LoadPercent")
	}
	
	// Optional: If absolutely NO signals are valid, we might return a special error state or just the empty snapshot.
	// But logic upstream handles partial data.

	return s
}
