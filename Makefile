

# temporal runs. Later we containerize with docker
run:
	reflex -r '\.go' -s -- sh -c "go run main.go"