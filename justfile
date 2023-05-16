run:
    cd ui && \
    npm run lint --fix && \
    npm run build && \
    cd ..
    go run cmd/twitch_predictions_recorder/main.go
