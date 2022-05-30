<h2>This is the Galton Board</h2>

<script lang="ts">
  const socket = new WebSocket('ws://localhost:8080/galton-board/');

  socket.addEventListener('open', (_) => {
    socket.send('Hello server!');
  });
</script>

<style lang="scss">
  h2 {
    color: #ff3e00;
  }
</style>
