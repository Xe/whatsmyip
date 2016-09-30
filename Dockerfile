FROM xena/go

ADD main.go /go/src/github.com/Xe/whatsmyip/main.go
RUN cd /go/src/github.com/Xe/whatsmyip \
 && go build github.com/Xe/whatsmyip \
 && go install
CMD /go/bin/whatsmyip
