# Demonaz

Simple Golang daemon starter

### Config Example
Create `config.json` like this:
```
{
  "daemons": [
    {
      "name": "test",
      "cmd": "php test.php",
      "enable": true,
      "worker_count": 2
    },
    {
      "name": "test2",
      "cmd": "php test2.php",
      "enable": true,
      "worker_count": 5
    }
  ]
}
```

### Start in dev mode
`go run config.json`
### Start in prod mode
`./demonaz config.json`
