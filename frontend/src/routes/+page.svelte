<script>
  import { goto } from "$app/navigation"

  function login() {
    goto("/login")
  }

  function signup() {
    goto("/signup")
  }
</script>

<div class="vcenter">
  <button on:click={login}>
    Log In
  </button>
  <button on:click={signup}>
    Sign Up
  </button>
</div>

<style>
  button {
    background-color: var(--darkergray);
    color: var(--white);
    display: flex;
    height: 100vh;
    width: 100%;
    justify-content: space-evenly;
    align-items: center;
  }
  button:hover {
    background-color: var(--darkestgray);
  }
  button:focus {
    background-color: var(--blackish);
  }
  .vcenter {
    font-size: 2em;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
  }
</style>