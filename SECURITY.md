# Security Policy

## 🔒 Supported Versions
| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | ✅ Yes              |
| < 1.0.0 | ❌ No               |

## 🐛 Reporting a Vulnerability
We take the security of `curl-imposter` seriously. If you discover a security vulnerability, please follow these steps:

1. **Do not** open a public GitHub issue or discussion.
2. Email your report securely to: `security@yourdomain.com` *(or use GitHub's private vulnerability reporting if enabled)*
3. Include:
   - Clear description of the vulnerability
   - Steps to reproduce
   - Affected version(s)
   - Potential impact
   - Suggested fix or mitigation (if applicable)

## ⏱️ Response Timeline
- **Acknowledge Receipt**: Within 48 hours
- **Initial Assessment**: Within 7 days
- **Patch & Release**: Within 14–30 days (depending on severity)
- **Public Disclosure**: Only after a fix is released and users have had reasonable time to update.

## 🔍 Security Architecture Notes
- `curl-imposter` runs entirely locally and **does not transmit data to third parties**.
- Mock data is stored locally in `~/.curl-imposter/mocks.json` with standard `0644` permissions.
- The parser uses `go-shellquote` to safely handle shell input, but always validate untrusted input when extending the CLI.
- Dependencies are kept minimal and updated via `go mod tidy` & `go get -u`.

## 🛡️ Responsible Disclosure
We believe in coordinated disclosure and will credit reporters who follow these guidelines (unless anonymity is requested). Thank you for helping keep `curl-imposter` secure! 🙏