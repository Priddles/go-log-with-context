# Reminder: First rule is the default target.
help: # Print this message.
help:
	@echo 'Make targets:'
	@egrep -h ":\s+# " $(MAKEFILE_LIST) | \
        sed -e 's/# //; s/^/    /' | \
        column -s: -t

check: # Format, lint, and test.
check:
	go fmt ./...
	golangci-lint run
	go test -p 1 -v $(TEST_ARGS) ./...
