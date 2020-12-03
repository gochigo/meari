# vim: set ts=4 sw=4 expandtab:
FROM scinix/golang

RUN \
    go get github.com/sio4/meari/cmd/meari ;\
    true

CMD /go/bin/meari
