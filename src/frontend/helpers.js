import Cookies from "js-cookie";
import { sendMessage } from "./socket";

export function logout(user, socket) {
  Cookies.remove("clapper-user");
  sendMessage(
    JSON.stringify({
      topic: "CLIENT_LEFT",
      user,
    }),
    socket
  );
  window.location.reload();
}
