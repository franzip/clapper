import { writable } from "svelte/store";
import { logout } from "./helpers";
import { initWs, sendMessage } from "./socket";
import { TOPICS } from "./constants";

const messageStore = writable([]);

let wsClient;

export default function initStore(user) {
  if (!user) return;

  wsClient = initWs(user);

  wsClient.addEventListener("open", () => {
    sendMessage(
      JSON.stringify({
        topic: TOPICS.CLIENT_JOINED,
        data: user,
      }),
      wsClient
    );
  });

  wsClient.addEventListener("message", (event) => {
    const { topic, data } = JSON.parse(event.data);
    messageStore.set(JSON.parse(data));
  });

  wsClient.addEventListener("close", () => logout(user, wsClient));

  return {
    subscribe: messageStore.subscribe,
  };
}
