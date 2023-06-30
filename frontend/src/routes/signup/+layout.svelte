<div class="container">
  <div class="lgap" />
  <div class="tgap" />

  <div class="area">
    <slot></slot>
  </div>

  <div class="rgap" />
  <div class="bgap" />
</div>

<style>
    .lgap {
    grid-area: lgap;
  }
  .tgap {
    grid-area: tgap;
  }
  .area {
    grid-area: login;
    background-color: var(--darkestgray);
    border-radius: .5em;
  }
  .rgap {

    grid-area: rgap;
  }
  .bgap {
    grid-area: bgap;
  }
  .container {
    height: 100vh;
    display: grid;
    grid-template-columns: 1fr 2fr 1fr;
    grid-template-rows: 1fr 1fr 1fr;
    grid-template-areas:
      "lgap  tgap  rgap "
      "lgap  login rgap "
      "lgap  bgap  rgap "
  }
</style>