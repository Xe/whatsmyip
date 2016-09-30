FROM xena/go

ADD . /go/src/github.com/Xe/whatsmyip
RUN cd /go/src/github.com/Xe/whatsmyip \
 && go build github.com/Xe/whatsmyip \
 && go install
CMD /go/bin/whatsmyip
