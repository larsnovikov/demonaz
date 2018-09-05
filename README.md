# Demonaz

Simple Golang daemon starter

### Config Example
Create `config.json` like this:
```
{
  "check_config_changes": false,
  "daemons": [
    {
      "name": "test",
      "cmd": "php test.php",
      "enable": true,
      "worker_count": 2,
      "log_file": "test2.log",
      "start_delay": 5
    },
    {
      "name": "test2",
      "cmd": "php test2.php",
      "enable": true,
      "worker_count": 5,
      "log_file": "test2.log",
      "start_delay": 5
    }
  ]
}
```

### Start in dev mode
`go run config.json`
### Start in prod mode
`./demonaz config.json`
### Some tips
If you want change config without manual restart, set `check_config_changes=true` in config and add demonaz to crontab(every minute) 
