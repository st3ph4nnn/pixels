COMPILER = go
OUTPUT = ./windows/pixels.exe
FILES = ../bin/.

default:
ifeq ($(OS), LINUX)
	$(eval OUTPUT=./linux/pixels)
	$(eval FILES=../bin/*)
	$(eval COMPILER=env GOOS=linux GOARCH=amd64 go)
endif
ifeq ($(RUN), RUN)
	$(COMPILER) run $(FILES)
else
	$(COMPILER) build -o $(OUTPUT) $(FILES)
endif