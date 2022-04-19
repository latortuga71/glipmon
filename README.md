# glipmon
Monitor New Clipboard Changes and print new contents to stdout.

## Best way to use
```
nohup ./gnip > /tmp/changes.txt &
keep your shell while its running.
```
## Why
* Could be useful, alot of people tend to copy and paste sensitive information often.

## Install
`go install github.com/latortuga71/glipmon/cmd/glipmon@latest`