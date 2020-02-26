## http-timeout

http-timeout restrict http server without timeout.
You must not use http server without timeout.

See [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/).

## Install

```bash
# MacOS 
curl -L https://github.com/wreulicke/http-timeout/releases/download/v0.0.1/http-timeout_0.0.1_darwin_amd64 -o /usr/local/bin/http-timeout

# Linux
curl -L https://github.com/wreulicke/http-timeout/releases/download/v0.0.1/http-timeout_0.0.1_linux_amd64 -o /usr/local/bin/http-timeout

# Windows
curl -L https://github.com/wreulicke/http-timeout/releases/download/v0.0.1/http-timeout_0.0.1_windows_amd64.exe -o <path-directory>/http-timeout.exe
```

## Usage


```
http-timeout ./...
```