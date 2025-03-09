# Run with Crontab

1. Build and place binary:
```bash
go build -o pgbackup
sudo cp pgbackup /usr/local/bin/
```

2. Place .env file:
```bash
sudo mkdir /etc/pgbackup
sudo cp .env /etc/pgbackup/
```

3. Open crontab:
```bash
crontab -e
```

4. Add one of these lines:
```bash
# Daily at 2 AM (cd is needed because .env file is in /etc/pgbackup)
0 2 * * * cd /etc/pgbackup && sudo /usr/local/bin/pgbackup

# Every 6 hours
# 0 */6 * * * cd /etc/pgbackup && sudo /usr/local/bin/pgbackup

# Weekly (Sunday 1 AM)
# 0 1 * * 0 cd /etc/pgbackup && sudo /usr/local/bin/pgbackup
```
