<script>
  import { onMount } from "svelte";
  import Cookies from "js-cookie";
  import { getWs } from "../socket";
  import { logout } from "../helpers";
  import Login from "./Login.svelte";
  import Clapper from "./Clapper.svelte";

  const user = Cookies.get("clapper-user");
  let wsClient = getWs();

  $: isLoggedIn = !!user;

  function handleLogout() {
    logout(user, wsClient);
  }
</script>

<main>
  <h1>Welcome to Clapper!👏👏👏</h1>
  {#if isLoggedIn}
    <Clapper {user} on:logout={handleLogout} />
  {:else}
    <Login />
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
