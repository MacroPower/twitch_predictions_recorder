run-all: run-ui run-api

run-ui:
    cd ui && \
    npm run lint --fix && \
    npm run build

run-api:
    go run cmd/twitch_predictions_recorder/main.go
