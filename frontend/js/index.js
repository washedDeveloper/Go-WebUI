const socket = new WebSocket("ws://localhost:3000/ws");

socket.onopen = function(e) {
   console.log("Socket connection established.");
   socket.send("Pong");
};

socket.onmessage = function(event) {
    console.log(`Data received from server: ${event.data}`);
};

socket.onclose = function(e) {
    console.log("Socket connected closed")
}