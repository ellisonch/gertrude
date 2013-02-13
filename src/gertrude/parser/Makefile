install: parser.go tokenizer.nn.go
	@go install

parser.go: parser.y
	@go tool yacc -o parser.go parser.y

tokenizer.nn.go: tokenizer.nex
	@nex tokenizer.nex
# @golex -t $< | gofmt > $@

test: install
	@go test

clean:
	@rm -f y.go y.output tokenizer.nn.go parser.go parser
