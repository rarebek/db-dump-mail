# Run with Crontab

1. Build and place binary:
```bash
go build -o pgbackup
sudo cp pgbackup /usr/local/bin/
```

2. Open crontab:
```bash
crontab -e
```

3. Add one of these lines:
```bash
# Daily at 2 AM
0 2 * * * cd /etc/pgbackup && /usr/local/bin/pgbackup

# Every 6 hours
# 0 */6 * * * cd /etc/pgbackup && /usr/local/bin/pgbackup

# Weekly (Sunday 1 AM)
# 0 1 * * 0 cd /etc/pgbackup && /usr/local/bin/pgbackup
```
