<h2>This is the Galton Board</h2>
<h3>The socket is currently: {socketOpen === 1 ? '🟢' : '🛑'}</h3>

<script lang="ts">
  let socketOpen: number;
  let reconnectInterval: NodeJS.Timer;

  const pollConnection = () => {
    if (socketOpen !== 1) {
      connectSocket();
    } else {
      console.log('already connected.');
    }
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
      if (!reconnectInterval) {
        reconnectInterval = setInterval(pollConnection, 1000);
      }

      socketOpen = socket.CLOSED;
    });
  };

  if (!reconnectInterval) {
    reconnectInterval = setInterval(pollConnection, 1000);
  }
</script>

<style lang="scss">
  h2,
  h3 {
    color: #ff3e00;
  }
</style>
