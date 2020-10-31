

# temporal runs. Later we containerize with docker
run_user:
	reflex -r '\.go' -s -- sh -c "go run main.go"

run_gateway:
	reflex -r '\.go' -s -- sh -c "go run gateway/server.go"