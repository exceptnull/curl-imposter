# Contributing to curl-imposter

Thank you for your interest in contributing to `curl-imposter`! This project is built to be lightweight, fast, and developer-friendly. We welcome contributions of all kinds: bug fixes, feature enhancements, documentation improvements, and performance optimizations.

## 🐛 Reporting Bugs
- Check existing [Issues](https://github.com/yourusername/curl-imposter/issues) to avoid duplicates.
- Use the provided bug report template.
- Include: OS, Go version, exact command run, expected vs actual output, and clear reproduction steps.

## 💡 Feature Requests
- Open an issue with the `enhancement` label.
- Clearly describe the use case, proposed behavior, and how it aligns with the project's **zero-config, single-binary** philosophy.

## 🛠️ Development Setup
1. Fork the repository and clone it locally
2. Ensure Go `1.20+` is installed (`go version`)
3. Install dependencies: `go mod tidy`
4. Build locally: `go build -o curl-imposter .`
5. Test the CLI: `./curl-imposter --help`

## 📝 Code Style & Standards
- Run `gofmt -w .` before committing
- Run `go vet ./...` and resolve all warnings
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use [Conventional Commits](https://www.conventionalcommits.org/) for commit messages
- Keep functions small, well-documented, and focused
- Add tests for new logic in `internal/` when applicable

## 🔀 Pull Request Process
1. Create a feature branch (`git checkout -b fix/issue-123`)
2. Commit changes with clear, descriptive messages
3. Push to your fork and open a PR against `main`
4. Fill out the PR template completely
5. Wait for review, address feedback, and ensure CI passes

## 📜 License
By contributing, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).

## 🙏 Thank You
Every contribution, no matter how small, helps make `curl-imposter` better for developers everywhere. We appreciate your time and effort!