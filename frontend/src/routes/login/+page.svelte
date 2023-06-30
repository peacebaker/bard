<script>

</script>

<form class="form">
  <div class="title">
    <h1>Login</h1>
  </div>

  <div class="username">
    Username
  </div>
  <input type="text">
  <div class="password">
    Password
  </div>
  <input type="password">

  <button class="submit" type="button">
    Submit
  </button>
</form>
 
<style>
  .title {
    text-align: center;
  }
  .form {
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    font-size: 1.25em;
    height: 100%;
    padding: 1em;
  }
  input {
    border: none;
    background-color: inherit;
    border-bottom: 1px solid var(--purple);
    color: inherit;
    font-size: .9em;
    padding: .25em;
    box-sizing: border-box;
  }
  button {
    font-size: 1em;
    margin-top: 1em;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0px;
  }
  .username {
    margin-top: 0px;
  }
  .password {
    margin-top: 12px;
  }
  .submit {
    align-self: center;
    background-color: var(--darkpurple);
    color: inherit;
    border: none;
    border-radius: .25em;
    padding: .25em .75em;
  }
</style>