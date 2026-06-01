# 🎭 curl-imposter

> Mock any API instantly from a single `curl` command. Capture real responses, replay them locally, and develop offline with zero configuration.

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Status](https://img.shields.io/badge/Status-Stable-brightgreen)](https://github.com/exceptnull/curl-imposter)

## ✨ Why `curl-imposter`?

Traditional API mocking tools (Postman, MockServer, WireMock) require complex setup, heavy dependencies, or manual JSON configuration. `curl-imposter` flips the workflow:
1. **Copy** a `curl` command from Chrome DevTools or your terminal
2. **Run** `curl-imposter learn "curl ..."`
3. **Serve** it locally with `curl-imposter serve`
4. **Develop** offline with exact, realistic API responses

Built in Go for speed, zero runtime dependencies, and single-binary distribution.

## 🚀 Features

- ⚡ **Zero-Config Mocking**: No YAML, no Docker, no complex DSLs
- 📦 **Exact Response Capture**: Status codes, headers, bodies, and latency
- ⏱️ **Realistic Latency Replay**: Simulates real API response times out of the box
- 🗄️ **Portable JSON Storage**: Human-readable, version-control friendly
- 🔒 **Safe Header Handling**: Auto-skips connection/content-length headers
- 🛡️ **Graceful Shutdown**: Clean `Ctrl+C` handling with in-flight request support
- 🐧 **Cross-Platform**: Single binary for Linux, macOS, Windows

## 📦 Installation

### From Source
```bash
git clone https://github.com/exceptnull/curl-imposter.git
cd curl-imposter
go build -o curl-imposter .
sudo mv curl-imposter /usr/local/bin/
```

### Via `go install` (once published)
```bash
go install github.com/exceptnull/curl-imposter@latest
```

## 🛠️ Usage

### 1. Capture a Real API Response
```bash
curl-imposter learn "curl -X GET https://api.example.com/v1/users -H 'Authorization: Bearer token123'"
```
**Output:**
```
🔍 Parsed: GET https://api.example.com/v1/users
📡 Hitting real API...
✅ Captured: 200 (142ms)
💾 Mock saved to ~/.curl-imposter/mocks.json
```

### 2. Start the Mock Server
```bash
curl-imposter serve
# Default: http://localhost:8080
# Custom port: curl-imposter serve --port 3000
```
**Output:**
```
🚀 Mock server running on http://localhost:8080
📊 Loaded endpoints:
   • GET /v1/users
```

### 3. Test Your Mock
```bash
curl http://localhost:8080/v1/users
# Returns the exact JSON captured earlier!
```

### 4. Clear & Reset
```bash
curl-imposter clear
# 🧹 Mock database cleared successfully.
```

## 🧠 How It Works

`curl-imposter` uses a streamlined, four-stage pipeline:

```
curl string → 🧩 Parser → 📡 Recorder → 💾 Storage → 🖥️ Mock Server
```

1. **Parser**: Uses `go-shellquote` to safely extract method, URL, headers, and body from raw shell strings
2. **Recorder**: Executes the request via Go's `net/http` with automatic gzip decompression & timeout handling
3. **Storage**: Persists structured `MockEntry` objects to `~/.curl-imposter/mocks.json`
4. **Server**: Builds an O(1) route lookup map (`METHOD /path` → `MockEntry`) and replays responses with realistic latency

## 📁 Project Structure

```
curl-imposter/
├── cmd/                 # CLI commands (learn, serve, clear)
├── internal/
│   ├── parser/          # curl string → structured request
│   ├── recorder/        # HTTP execution & response capture
│   ├── storage/         # JSON file persistence
│   └── server/          # HTTP mock router & handler
├── main.go              # Entry point
└── go.mod               # Dependencies
```

## 🤝 Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure your code passes `go vet` and `go fmt`. See `CONTRIBUTING.md` for detailed guidelines.

## 📜 License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for more information.

## 🙏 Acknowledgments

- Inspired by the need for lightweight, developer-first API mocking
- Built with [Cobra](https://github.com/spf13/cobra) and [go-shellquote](https://github.com/kballard/go-shellquote)
- Designed for developers who live in the terminal