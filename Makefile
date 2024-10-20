.PHONY: frontend
# 变量
TARGET = target/app
APP = mine
BACK_SOURCE = cmd/main.go
FRONT_SOURCE = frontend/mine-fe

# all
start:
	@make backend
	@make frontend

stop:
	@pgrep $(APP) | xargs kill -SIGINT
	@pgrep node  | xargs kill -SIGINT

# 后端
clean:
	@go clean
	@rm -rf $(TARGET)
build: clean
	@go build -o $(TARGET)/$(APP) $(BACK_SOURCE)
backend: build
	@cd $(TARGET) && ./$(APP) -config=../../etc/config.yaml &
# 前端
frontend:
	@cd $(FRONT_SOURCE) && npm run dev &