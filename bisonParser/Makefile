gertrude-parser: parser.cpp tokens.cpp main.cpp parser.hpp node.h
	@g++ -Wall -o gertrude-parser parser.cpp tokens.cpp main.cpp

parser.cpp parser.hpp: parser.y
	@bison -d -o parser.cpp parser.y

tokens.cpp: tokens.l
	@flex -o tokens.cpp tokens.l

test: parser sample.grt
	@./gertrude-parser < sample.grt

clean:
	@rm -f *.exe *.o tokens.cpp parser.cpp parser.hpp
