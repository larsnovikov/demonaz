<p align="center">
    <img src="https://www.metal-archives.com/images/2/8/7/287_artist.jpg?4821" height="100px">
    <h1 align="center">Demonaz</h1>
    <br>
</p>

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
      "worker_count": 2,
      "log_file": "test2.log",
    },
    {
      "name": "test2",
      "cmd": "php test2.php",
      "enable": true,
      "worker_count": 5,
      "log_file": "test2.log",
    }
  ]
}
```

### Start in dev mode
`go run config.json`
### Start in prod mode
`./demonaz config.json`
