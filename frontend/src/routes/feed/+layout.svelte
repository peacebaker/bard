<script>

</script>

<div class="container">

  <nav class="mainNav">
    <div class="left" >
      <a href="/">
        LowKey Social
      </a>
    </div>
    
    <div class="right">
      <a href="/login">Login</a>
      |
      <a href="/logout">Logout</a>
    </div>
  </nav>

  <nav class="leftNav">
    <ul>
      <li>
        Chats
      </li>
      <li>
        Posts
      </li>
    </ul>
  </nav>

  <div class="content">
    <slot></slot>
  </div>
</div>


<style>

  /*
  * thinking... dark mode and... darker mode?
  * tryin to keep this low key, after all
  */

  @import url("https://code.cdn.mozilla.net/fonts/fira.css");

  .mainNav {
    padding: 3px;
    grid-area: mainNav;
    font-size: 2em;
  }
  .mainNav .left {
    float: left;
  }
  .mainNav .right {
    float: right;
  }
  .leftNav {
    grid-area: leftNav;
  }
  .content {
    grid-area: content;
  }
  .container {
    display: grid;
    grid-template-areas:
      "mainNav mainNav mainNav mainNav mainNav mainNav"
      "leftNav content content content content content";

    font-family: "Fira Sans";
    font-size: 1em;
    
  }
</style>