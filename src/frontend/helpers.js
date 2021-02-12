import Cookies from "js-cookie";
import { sendMessage } from "./socket";
import { TOPICS } from "./constants";

export function logout(user, socket) {
  Cookies.remove("clapper-user");
  sendMessage(
    JSON.stringify({
      topic: TOPICS.CLIENT_LEFT,
      data: user,
    }),
    socket
  );
  window.location.reload();
}
