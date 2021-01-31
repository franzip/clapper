import { writable } from "svelte/store";
import { logout } from "./helpers";
import { initWs, sendMessage } from "./socket";

const messageStore = writable("");

let wsClient;

export default function initStore(user) {
  if (!user) return;

  wsClient = initWs(user);

  wsClient.addEventListener("open", () => {
    sendMessage(
      JSON.stringify({
        topic: "CLIENT_JOINED",
        user,
      }),
      wsClient
    );
  });

  wsClient.addEventListener("message", function (event) {
    console.log(event);
    messageStore.set(event.data);
  });

  wsClient.addEventListener("close", () => logout(user, wsClient));

  return {
    subscribe: messageStore.subscribe,
  };
}
