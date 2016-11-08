function connect() {
  ws = new WebSocket("ws://localhost:3000/socket");

  ws.onopen = function() {
    el = document.getElementById("open-connection");
    el.style.display = "block";
    el = document.getElementById("closed-connection");
    el.style.display = "none";
  };

  ws.onmessage = function(evt) {
    msg = JSON.parse(evt.data);
    el = document.getElementById("messages");
    el.innerHTML += ("<li>" + msg.formatted + " (" + msg.received + ")" + "</li>");
  };

  ws.onclose = function() {
    el = document.getElementById("open-connection");
    el.style.display = "none";
    el = document.getElementById("closed-connection");
    el.style.display = "block";

    setTimeout(connect, 5000);
  };
}

connect();

function sendMessage(e) {
  el = document.getElementById("message");
  val = el.value;
  ws.send(JSON.stringify(val));
  el.value = "";
  el.focus();
  return false;
}
