<script>
  import { onMount, onDestroy } from "svelte";
  import initStore from "../store";
  import { createEventDispatcher } from "svelte";

  let unsubscribe;
  let clients = [];

  export let user;

  const dispatch = createEventDispatcher();

  function logout() {
    dispatch("logout");
  }

  const store = initStore(user);
  unsubscribe = store.subscribe((newClients) => {
    clients = newClients;
  });

  onDestroy(unsubscribe);
</script>

<p>Connected as {user}</p>
<h4>Choose the Clapper MC:</h4>
<div class="mc-choice">
  {#each clients as { id, username }}
    <button {id}>{username}</button>
  {/each}
</div>

<div class="commands">
  <button on:click={logout}>Logout</button>
</div>
