install: parser.go tokenizer.go
	@go install

parser.go: parser.y
	@go tool yacc -o parser.go parser.y

tokenizer.go: tokenizer.l
	@golex -t $< | gofmt > $@

test: install
	@go test

clean:
	@rm -f y.go y.output tokenizer.go parser.go parser
