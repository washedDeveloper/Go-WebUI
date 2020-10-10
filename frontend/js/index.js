const socket = new WebSocket("ws://localhost:1234");

socket.onopen = function(e) {
   alert("Socket connection established.");
};

socket.onmessage = function(event) {
    console.log(`Data received from server: ${event.data}`);
    socket.send("Pong");
};

socket.onclose = function(e) {
    console.log("Socket connected closed")
}