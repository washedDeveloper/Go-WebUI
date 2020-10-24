# Go Web UI

Go Web UI is an example of using Chrome's app flag to serve a desktop application from Go.

## Inspiration
I wanted to create a way to give a Go program a UI, without the bloat that comes with it, like Electron. This uses the system installed Chrome, so the program *basically* has no UI dependencies.

## Testing

Install `github.com/gorilla/mux` & `github.com/gorilla/websocket` and run:

```bash
go run main.go
```

## IPC
For IPC, I've used web sockets from the main Go process, to the JavaScript renderer. This is currently a POC, and will be changed in future to perform close to Electron.


## Inspiration
[MIT](https://choosealicense.com/licenses/mit/)
