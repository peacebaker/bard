<script>
  export let showErr = false;
  export let errName = "";
  export let errMessage = "";
  export let errSoft = false;

  function close() {
    showErr = false;
  }
</script>

<dialog open>
  {#if errSoft}
    <h2>Failure</h2>
  {:else}
    <h2>Error</h2>
  {/if}
  <p>{errName}</p>
  <p>{JSON.stringify(errMessage)}</p>

  <button on:click={close}>
    Close
  </button>
</dialog>
