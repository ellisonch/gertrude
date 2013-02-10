parser: parser.go tokenizer.go
	@go build -o parser parser.go tokenizer.go

parser.go: parser.y
	@go tool yacc -o parser.go parser.y

tokenizer.go: tokenizer.l
	@golex -t $< | gofmt > $@

test: parser sample.grt
	@./parser < sample.grt

clean:
	@rm -f *.exe *.o tokens.cpp parser.cpp parser.hpp
