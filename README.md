# twitch_predictions_recorder

Records twitch prediction data for analysis.

You could use your findings with [Twitch Channel Points Miner][tcpm].

## Installation

You can download a [release][releases], pull the [docker image][docker-hub], or
build it yourself:

```bash
go install github.com/MacroPower/twitch_predictions_recorder/cmd/twitch_predictions_recorder@main
```

## Usage

```
Usage: twitch_predictions_recorder

Flags:
  -h, --help                               Show context-sensitive help.
      --twitch-client-id=STRING            Twitch Client ID ($TWITCH_CLIENT_ID).
      --twitch-secret=STRING               Twitch Secret ($TWITCH_SECRET).
      --streamers-file="streamers.txt"     List of streamers to monitor ($STREAMERS_FILE).
      --db.type="postgres"                 Database type. One of: [postgres, test] ($DB_TYPE)
      --db.pg.host="info"                  PG Host ($DB_PG_HOST).
      --db.pg.port=5432                    PG Port ($DB_PG_PORT).
      --db.pg.ssl-mode="prefer"            PG SSL Mode ($DB_PG_SSL_MODE).
      --db.pg.user=STRING                  PG User ($DB_PG_USER).
      --db.pg.password=STRING              PG Password ($DB_PG_PASSWORD).
      --db.pg.db-name="postgres"           PG DB Name ($DB_PG_DB_NAME).
      --db.time-zone="America/New_York"    Time zone name ($DB_TIME_ZONE).
      --metrics.disable                    Disable metrics ($METRICS_DISABLE).
      --metrics.path="/metrics"            Path to serve metrics on ($METRICS_PATH).
      --metrics.address=":8080"            Address to serve metrics on ($METRICS_ADDRESS).
      --metrics.timeout=60s                HTTP timeout ($METRICS_TIMEOUT).
      --log.level="info"                   Log level ($LOG_LEVEL).
      --log.format="logfmt"                Log format. One of: [logfmt, json] ($LOG_FORMAT)
```

### Getting a Client ID & Secret

Go to [dev.twitch.tv/console][twitch-console] and register a new application.

You can use a localhost redirect URI.

Once created, you can see your Client ID and generate a secret.

## Notes on betting

### Automation is not always better

Consider a prediction on the outcome of a competitive game. You notice that
Surefour is on the enemy team. However, while you know Surefour is one of the
best gamers alive and is very likely to win, almost nobody else in chat knows
who he is. Thus the chat predicts that the probability of the steamer winning
is 55%.

In this case, your strategy may also bet along with the highest probability
outcome, if you have identified that the chat typically overestimates the
probability of the less likely outcome. However, if you were the one placing the
bet, you know that the probability of winning is actually around 30%. In which
case, you would bet against the strategy, given your additional knowledge of the
situation.

### Be careful with bet sizing

Make sure to consider whether or not a certain return is possible after making
your bet. e.g. if the pool was 5,000 total points, placing a 10,000 point bet
would drastically impact the odds you saw when placing the bet.

### Slippage

There is some slippage due to Twitch rounding all points down when rewarding
points. This means that you will need to bet large enough amounts that this is
minimized. For example, if you only bet 10 points, you can lose a significant
percentage of your reward to slippage. If you bet a few thousand points however,
you will only lose a fraction of a percent of your reward.

### Bet timing

An additional area of concern is when you place a bet. You will likely place
bets slightly before the timer has expired, leaving a short time for additional
users to place bets which influence the overall odds. If you bet earlier (e.g.
if the streamer typically closes bets early), this can be troublesome. It is
possible to account for this by recording how the odds change over time, however
this is not yet implemented.

### Edge cases

If you always bet the same or very similar amounts regardless of outcome
probability, you can sometimes run into cases which seem strange.

Two opposing strategies can both win.

| Bet       | Winner | Loser | Greatest Probability | Least Probability |
| --------- | ------ | ----- | -------------------- | ----------------- |
| 1000      | 10%    | 90%   | -1000                | +9000             |
| 1000      | 55%    | 45%   | +550                 | -1000             |
| 1000      | 55%    | 45%   | +550                 | -1000             |
| **TOTAL** | -      | -     | +100                 | +7000             |

The money is slurped from the relatively large number of users on bet #1,
because you didn't bet less when the risk was greater.

Two opposing strategies can both lose.

| Bet       | Winner | Loser | Greatest Probability | Least Probability |
| --------- | ------ | ----- | -------------------- | ----------------- |
| 1000      | 90%    | 10%   | +100                 | -1000             |
| 1000      | 90%    | 10%   | +100                 | -1000             |
| 1000      | 45%    | 55%   | -1000                | +550              |
| 1000      | 45%    | 55%   | -1000                | +550              |
| **TOTAL** | -      | -     | -1800                | -900              |

[releases]: https://github.com/MacroPower/twitch_predictions_recorder/releases
[docker-hub]: https://hub.docker.com/r/macropower/twitch_predictions_recorder
[twitch-console]: https://dev.twitch.tv/console/apps/create
[tcpm]: https://github.com/Tkd-Alex/Twitch-Channel-Points-Miner-v2
