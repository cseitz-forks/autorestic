# Contributing

## Running

```bash
docker build -t autorestic .
# get shell
docker run -it autorestic bash
# create config
mkdir -p /root/.config
```

```yaml filename="/root/.config/.autorestic.yml"
version: 2
 
locations:
  home:
    from: /root/data-1
    to: local-1
    cron: '*/5 * * * *'
    options:
      forget:
        keep-last: 5 # always keep at least 5 snapshots
        keep-hourly: 3 # keep 3 last hourly snapshots
        keep-daily: 4 # keep 4 last daily snapshots
        keep-weekly: 1 # keep 1 last weekly snapshots
        keep-monthly: 12 # keep 12 last monthly snapshots
        keep-yearly: 7 # keep 7 last yearly snapshots
        keep-within: '14d' # keep snapshots from the last 14 days
 
backends:
  local-1:
    type: local
    path: '/data/backups/local-1'
    key: 'password'

```


```bash
# run cron backup only if we are ready to do so
autorestic cron --dry-run | grep -q "Cron is ready" && autorestic cron
```

```bash
# start autorestic rundeck job only if we have crons that are due
autorestic cron --dry-run | grep -q "Cron is ready" && rd run --id [JOB-ID]
# add this to crontab and then rundeck job will only run when we have things to output.
# this guarentees every log in rundeck activity will actually be useful for providing info on backups
```


```bash
restic -r /data/backups/local-1 snapshots
```

