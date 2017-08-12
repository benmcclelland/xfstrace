# xfstrace
a test example of tracing `XFS_IOC_FSCOUNTS` ioctl
[article explaining this code](https://medium.com/@ben.mcclelland/ioctl-trace-in-40-more-lines-of-go-3131848fadc5)

# getting:
`go get github.com/benmcclelland/xfstrace`

# building:
`GOOS=linux go build -o xfstrace  *.go`
