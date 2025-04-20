.PHONY: test
test:
ifdef SERVICE
	@echo "Running tests for $(SERVICE)..."
	go test -v ./$(SERVICE)/...
else
	@echo "Running tests for all services..."
	go test -v ./api_gateway/...
	go test -v ./appointment_service/...
endif

.PHONY: up-test-env
up-test-env:
	@echo "Starting test environment with docker-compose..."
	@docker-compose up --build -d

# Проверка, поднят ли docker-compose
.PHONY: ensure-up
ensure-up:
	@echo "Checking if API Gateway is up..."
	@if ! curl -s --fail $${API_GATEWAY_ADDR:-http://localhost:8080}/healthz >/dev/null; then \
		echo "Not running — starting docker-compose..."; \
		docker-compose up --build -d; \
		sleep 5; \
	else \
		echo "Docker-compose environment is already up."; \
	fi

.PHONY: integration-test
integration-test: ensure-up
	@echo "Running integration tests..."
	go test -v -tags=integration ./tests/integration

.PHONY: clean
clean:
	docker-compose down
	@echo "Cleaned up docker-compose environment."
