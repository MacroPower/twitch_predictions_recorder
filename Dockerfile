FROM golang:1.17

COPY twitch_predictions_recorder /usr/local/bin/

ENTRYPOINT ["twitch_predictions_recorder"]
