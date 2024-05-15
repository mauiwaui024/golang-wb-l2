CC=go

all: my_wget

rebuild: clean all

my_wget:
	$(CC) build -o my_wget main.go 
clean:
	rm -rf my_wget
download:
	./my_wget get-visas.ru
