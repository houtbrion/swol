BINDIR=/usr/local/bin
SRC=swol.go
BIN=swol
GOOUT=go.mod
CP=cp

all: ${SRC}
	go mod init main
	go build ${SRC}

install:
	${CP} ${BIN} ${BINDIR}

clean:
	${RM} ${BIN} ${GOOUT}
