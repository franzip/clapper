<script>
  import Cookies from "js-cookie";

  function handleSubmit(event) {
    const user = new FormData(event.target).get("user");
    if (user.length >= 4) {
      Cookies.set("clapper-user", user);
      window.location.reload();
    }
  }
</script>

<p>Please enter your name:</p>
<form on:submit|preventDefault={handleSubmit}>
  <input id="user" name="user" minlength="4" type="text" />
  <button>Connect</button>
</form>
