let wsClient;

export function initWs(user) {
  wsClient = new WebSocket(`ws://${document.location.host}/ws?user=${user}`);

  return wsClient;
}

export function getWs() {
  return wsClient;
}

export function sendMessage(message) {
  if (wsClient.readyState <= 1) {
    wsClient.send(message);
  }
}
