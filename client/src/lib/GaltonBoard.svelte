<h2>This is the Galton Board</h2>
<h3>The socket is currently: {socketOpen === 1 ? '🟢' : '🛑'}</h3>

<script lang="ts">
  let socketOpen: number;
  let reconnectInterval: NodeJS.Timer;

  const pollConnection = () => {
    if (reconnectInterval) {
      return;
    }

    reconnectInterval = setInterval(() => {
      if (socketOpen !== 1) {
        connectSocket();
      } else {
        console.log('already connected.');
      }
    }, 1000);
  };

  const connectSocket = () => {
    let socket = new WebSocket('ws://localhost:8080/galton-board/');

    socket.addEventListener('open', (_) => {
      clearInterval(reconnectInterval);
      reconnectInterval = null;

      socket.send('Hello server!');
      socketOpen = socket.readyState;
    });

    socket.addEventListener('close', (_) => {
      pollConnection();

      socketOpen = socket.CLOSED;
    });
  };

  pollConnection();
</script>

<style lang="scss">
  h2,
  h3 {
    color: #ff3e00;
  }
</style>
