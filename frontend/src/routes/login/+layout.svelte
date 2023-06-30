<div class="container">
  <div class="lgap" />
  <div class="tgap" />

  <div class="area">
    <slot />
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
  border-radius: .75em;
}
.rgap {
  grid-area: rgap;
}
.bgap {
  grid-area: bgap;
}
.container {
  display: grid;
  height: 100vh;
  grid-template-columns: 3fr 1fr 3fr;
  grid-template-rows: .9fr 1fr .9fr;
  grid-template-areas:
    "lgap  tgap  rgap "
    "lgap  login rgap "
    "lgap  bgap  rgap "
}
</style>