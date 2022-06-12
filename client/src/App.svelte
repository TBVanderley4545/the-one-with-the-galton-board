<main>
  <img src={logo} alt="Svelte Logo" />
  <GaltonBoard />
</main>

<script lang="ts">
  import logo from './assets/svelte.png';
  import GaltonBoard from './lib/GaltonBoard.svelte';
</script>

<style lang="scss">
  :root {
    background-color: #1c1b1a;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
      'Open Sans', 'Helvetica Neue', sans-serif;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
  }

  img {
    height: 16rem;
    width: 16rem;
  }
</style>
