<script>
  import config from "$lib/config.js"
  import { CookieJar } from "$manual/CookieJar/CookieJar.js";
  import ShowErr from "$lib/ShowErr.svelte";

  let username = "";
  let email = "";
  let password = "";
  let confirm = "";

  let firstName = "";
  let lastName = "";
  let nickName = "";

  let showErr = false;
  let errName = "";
  let errMessage = "";
  let errSoft = false;

  async function signMeUp() {

    // encode the form into JSON
    let formData = {
      neighborhood: "Alpha",
      username: username,
      email: email,
      password: password,
      firstName: firstName,
      lastName: lastName,
      nickName: nickName,
    }
    let form = JSON.stringify(formData)
    

    // send request to the backend server
    let response = await fetch(`${config.Backend}/signup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: form
    })
    
    // wrangle the response
    let data = await response.json();
    
    // check for and display an error
    if (data.Err) {
      console.log("butts")
      errName = data.Err.Code;
      errMessage = data.Err.Message;
      errSoft = (response.status === 200);
      showErr = true;
      return
    }

    // save those cookies
    let session = data.Session;
    CookieJar.set("neighborhood", session.neighborhood)
    CookieJar.set("username", session.username)
    CookieJar.set("token", session.token)
  }
</script>

{#if showErr}
  <ShowErr bind:showErr={showErr} {errName} {errMessage} {errSoft} />
 {/if}

<div class="title">
  <h1>Sign Up (Alpha)</h1>
</div>

<form>

  <fieldset>
    <legend>Login</legend>

    <div class="username">
      Username<br>
      <input type="text" bind:value={username}><br>
      Email<br>
      <input type="text" bind:value={email}><br>
    </div>

    <div class="password">
      Password<br>
      <input type="password" bind:value={password}><br>
      Confirm Password<br>
      <input type="password" bind:value={confirm}><br>
    </div>

  </fieldset>

  <fieldset>
    <legend>Names (These don't have to be legal or anything.)</legend>
    First Name<br>
    <input type="text" bind:value={firstName}><br>
    Last Name<br>
    <input type="text" bind:value={lastName}><br>
    Nickname<br>
    <input type="text" bind:value={nickName}><br>
  </fieldset>
    
  <fieldset>
    <legend>Credit Card</legend>
    Not implemented yet.
    Lucky you, alpha testers!  xP
  </fieldset>

  <div class="captcha">
    <!-- definitely need one of these before going live, lol -->
  </div>

  <!-- submit button goes here -->
  <button class="submit" type="button" on:click|once={signMeUp}>
    Submit
  </button>

</form>

<style>

  @import url("https://code.cdn.mozilla.net/fonts/fira.css");

  .title {
    text-align: center;
  }
  form {
    font-size: 1.25em;
    padding: .5em;
  }
  fieldset {
    border-color: var(--darkblue);
    border-radius: .5em;
    margin-bottom: .5em; 
    padding: .75em .5em;
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
  .username {
    margin-top: 0px;
    float: left;
  }
  .password {
    margin-top: 0px;
    float: right;
  }
  .submit {
    display: block;
    align-self: center;
    background-color: var(--darkpurple);
    border: none;
    border-radius: .25em;
    color: inherit;
    font-size: 1em;
    margin-top: 1.5em;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: .75em;
    padding: .25em .75em;
  }
</style>