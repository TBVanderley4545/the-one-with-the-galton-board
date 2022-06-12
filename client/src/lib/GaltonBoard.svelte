<h2>This is the Galton Board</h2>
<h3>The socket is currently: {socketOpen === 1 ? '🟢' : '🛑'}</h3>
<p>Message from the socket: {socketMessage}</p>

<form on:submit|preventDefault={handleFormSubmit}>
  <input id="quantity" type="number" />
  <label for="quantity">Enter the number of levels to have for the board.</label>
  <button type="submit">Submit</button>
</form>

<script lang="ts">
  import type { GaltonBoard } from './types/GaltonBoard';
  import { SocketMessage, ClientMessages } from './utilities/sockets';

  let socket: WebSocket;
  let socketOpen: number;
  let reconnectInterval: NodeJS.Timer;
  let socketMessage: string;
  let board: GaltonBoard;

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
    socket = new WebSocket('ws://localhost:8080/galton-board/');

    socket.addEventListener('open', (_) => {
      clearInterval(reconnectInterval);
      reconnectInterval = null;

      const message: SocketMessage = {
        msg: ClientMessages.Opened,
      };

      socket.send(JSON.stringify(message));

      socketOpen = socket.readyState;
    });

    socket.addEventListener('message', (ev) => {
      socketMessage = ev.data;
    });

    socket.addEventListener('close', (_) => {
      pollConnection();

      socketOpen = socket.CLOSED;
    });
  };

  const handleFormSubmit = (e: SubmitEvent) => {
    const submittedFormElm = <HTMLFormElement>e.currentTarget;
    const boardDepth = submittedFormElm.elements['quantity'].value;

    const message: SocketMessage = {
      msg: ClientMessages.CreateBoard,
      quantity: Number.parseInt(boardDepth),
    };

    socket.send(JSON.stringify(message));
  };

  pollConnection();
</script>

<style lang="scss">
  h2,
  h3 {
    color: #ff3e00;
  }

  p {
    color: #aaaaaa;
  }
</style>
