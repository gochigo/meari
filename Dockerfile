# vim: set ts=4 sw=4 expandtab:
FROM scinix/golang

RUN \
    go get github.com/gochigo/meari/cmd/httpd ;\
    true

CMD /go/bin/httpd
