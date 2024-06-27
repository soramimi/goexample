run: cat
	go run main.go

build: cat
	go build main.go

cat:
	g++ -c cat.cpp

clean:
	rm -f main
	rm -f *.o
