<script>
  import { onMount, onDestroy } from "svelte";
  import initStore from "../store";
  import { createEventDispatcher } from "svelte";

  let unsubscribe;
  let messages = [];

  export let user;

  const dispatch = createEventDispatcher();

  function logout() {
    dispatch("logout");
  }

  const store = initStore(user);
  unsubscribe = store.subscribe((currentMessage) => {
    messages = [...messages, currentMessage];
  });

  onDestroy(unsubscribe);
</script>

<h4>Choose the Clapper MC:</h4>
<p>Connected as {user}</p>
<div class="mc-choice">
  <button>Apple</button>
  <button>Samsung</button>
  <button>Sony</button>
</div>

<div class="commands">
  <button on:click={logout}>Logout</button>
</div>
