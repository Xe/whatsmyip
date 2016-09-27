FROM xena/go

ADD main.go /go/src/github.com/Xe/whatsmyip/main.go
RUN go build github.com/Xe/whatsmyip
CMD /go/whatsmyip
