<slot></slot>

<style>
  @import url("https://code.cdn.mozilla.net/fonts/fira.css");
  @import '/node_modules/normalize.css/normalize.css';

  :global(body) {
    background-color: var(--blackish);
    color: var(--white);
    font-family: "Fira Sans";

    --slightlydarkerblack: hsl(0, 0%, 0%);
    --black: hsl(0, 0%, 5%);
    --blackish: hsl(0, 0%, 10%);
    --darkestgray: hsl(0, 0%, 20%);
    --darkergray: hsl(0, 0%, 30%);
    --darkgray: hsl(0, 0%, 40%);
    --gray: hsl(0, 0%, 50%);
    --lightgray: hsl(0, 0%, 60%);
    --lightergray: hsl(0, 0%, 70%);
    --lightestgray: hsl(0, 0%, 80%);
    --whitish: hsl(0, 0%, 90%);
    --white: hsl(0, 0%, 95%);
    --slightlybrighterwhite: hsl(0, 0%, 100%);

    --darkpurple: hsl(270, 80%, 40%);
    --purple: hsl(270, 100%, 60%);
    
    --darkblue: hsl(205, 80%, 40%);
    --blue: hsl(205, 100%, 60%);

    --darkaqua: hsl(180, 80%, 40%);
    --aqua: hsl(180, 100%, 60%);

    --darkgreen: hsl(160, 80%, 40%);
    --green: hsl(160, 100%, 60%);

    --darkpink: hsl(310, 80%, 40%);
    --pink: hsl(310, 100%, 60%);
  }
</style>