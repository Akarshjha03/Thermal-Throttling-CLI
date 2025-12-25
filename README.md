# Thermal Throttling Analyzer

![Project Banner](assets/banner.svg)

![CLI](https://img.shields.io/badge/CLI-Cobra-green)
![Platform](https://img.shields.io/badge/platform-Windows-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Go](https://img.shields.io/badge/go-1.21-blue?logo=go)

**Thermal Throttling Analyzer (TTA)** is a powerful, lightweight, and Windows-only CLI diagnostic tool built in Go. It empowers developers, gamers, and power users to detect, analyze, and explain CPU thermal throttling events in real-time.

---

## 🚀 Project Idea & Use Cases

Modern CPUs invoke thermal throttling to protect themselves from overheating, but this often results in unexplained performance drops (FPS lag, compile time slowdowns). **TTA** bridges the gap between raw sensor data and actionable insights.

### Common Use Cases:
- **Gamers**: Diagnose why your frame rate suddenly drops after 30 minutes of gameplay.
- **Developers**: Understand if thermal limits are affecting your compile times or render jobs.
- **Overclockers**: Verify system stability and cooling efficiency under load.
- **SysAdmins**: Quick health checks on Windows servers without heavy GUI tools.

---

## ✨ Key Features

- **🔥 Real-time Monitoring**: Watch CPU state, temperature, and frequency live.
- **📊 Historical Analysis**: Analyze past logs to correlate slowdowns with thermal events.
- **🏥 Doctor Mode**: One-command system health check.
- **📉 Throttling Detection**: Smart state machine detects when and why throttling occurs.
- **🎈 Lightweight**: Zero dependencies, single binary executable.

---

## 📦 Installation

You can build the project from source (requires Go 1.21+):

```bash
git clone https://github.com/Akarshjha03/ThermalThrottlingAnalyzer.git
cd ThermalThrottlingAnalyzer
go mod tidy
go build -o tta.exe ./cmd/tta
```

Or install directly via `go install`:

```bash
go install github.com/Akarshjha03/ThermalThrottlingAnalyzer/cmd/tta@latest
```

---

## 🛠️ Commands

### 1. Check System Status
Quickly verify if your system is currently throttling or healthy.
```bash
tta status
```
*Output: "Thermal State: NORMAL" or "Thermal State: THROTTLING"*

### 2. Live Monitoring
Watch your CPU metrics update in real-time. perfect for keeping open on a second monitor while testing.
```bash
tta watch
```

### 3. Analyze History
Investigate what happened in the last few hours.
```bash
tta analyze --last 2h
```

### 4. Health Check (Doctor)
Run a comprehensive diagnostic of your thermal sensors and configuration.
```bash
tta doctor
```

### 5. View Logs
Dump raw event logs for external processing.
```bash
tta log --today
```

---

## 🤝 Contributions

Contributions are always welcome! We'd love to have your help to make TTA better.

1.  **Fork** the repository.
2.  **Create** a feature branch (`git checkout -b feature/NewSensor`).
3.  **Commit** your changes (`git commit -m 'Add new sensor support'`).
4.  **Push** to the branch (`git push origin feature/NewSensor`).
5.  **Open** a Pull Request.

Please ensure you run `go fmt ./...` before submitting!

---

## ⭐ Support

If you find this project useful, please consider **leaving a star** on GitHub! It helps us know you appreciate the work and motivaes further development.

Made with ❤️ in Go.
