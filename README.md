# glipmon
Monitor New Clipboard Changes and print new contents to stdout.

## Best way to use
```
nohup ./glipmon > /tmp/changes.txt &
if you have a standard reverse shell this will just write to the shell stdout not the file
```
## Why
* Could be useful, alot of people tend to copy and paste sensitive information often.

## Install
`go install github.com/latortuga71/glipmon/cmd/glipmon@latest`
