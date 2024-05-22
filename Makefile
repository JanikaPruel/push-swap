.PHONY: all clean push_swap checker

all: push_swap checker

push_swap:
	go build -o bin/push_swap ./cmd/push_swap

checker:
	go build -o bin/checker ./cmd/checker

clean:
	rm -rf bin

run_push_swap:
	./bin/push_swap $(ARGS)

run_checker:
	./bin/checker $(ARGS)
