# local-proxy

Could be used to mock entrypoint response by config.

## Usage

### Command-Line Parameters

- -h  
  ```To show help.```
- -config.path  
  ```To specify config file path(should be json file). Default is "./proxy-config.json".```
- -listen.port  
  ```To specify listen port of service. Default is 8080.```

### Http Management Entrypoint

Request Method: any

- /-/list   
  ```To show proxy configs current using.```
- /-/reload  
  ```To reload proxy configs without restarting program.```

### Proxy Config

```json
[
  {
    "mockUrl": "/test",
    "requestMethod": "POST",
    "mockResult": "{\"code\":200,\"message\":\"test invoked successfully\"}"
  }
]
```

- mockUrl  
  ```Path you want to mock, without proctol and port.```
- requestMethod  
  ```Request method of your entrypoint, ignore case.```
- mockResult  
  ```Expect entrypoint response value, should be json.```