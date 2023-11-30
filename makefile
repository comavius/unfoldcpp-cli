NAME = "unfoldcpp"

install:
	@echo "Installing..."
	@go build -o $(NAME) main.go
	chmod +x $(NAME)
	@echo "Done!"
