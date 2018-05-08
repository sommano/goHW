all: check-gofmt
 
check-gofmt:
@if [ -n "$(shell go fmt -1 .)" ] then \
	echo 1>&2 'The following files need to be formatted:'; \
	gofmt -1 .; \
	exit 1; \
fi