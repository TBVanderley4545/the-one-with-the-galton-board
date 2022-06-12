<h3>The socket is currently: {socketOpen === 1 ? '🟢' : '🛑'}</h3>

{#if !board || !board.Columns || board.Columns.length === 0}
  <form class="board-create-form" on:submit|preventDefault={handleFormSubmit}>
    <label for="quantity">Enter the number of levels to have for the board:</label>
    <input id="quantity" type="number" />
    <button type="submit">Submit</button>
  </form>
{:else}
  <div class="galton-board__wrap">
    <h2>Galton Board</h2>

    <form on:submit|preventDefault={handleAddBalls}>
      <label for="quantity">Enter the number of balls to add:</label>
      <input id="quantity" type="number" />
      <button type="submit">Submit</button>
    </form>

    <button class="galton-board__reset" on:click={handleResetBoard}>Reset board</button>

    <p>Ball count: {ballCount}</p>

    <div class="galton-board">
      {#each board.Columns as column, idx}
        <div class="galton-board__column-wrap">
          <div class="galton-board__column">
            <span>
              {column?.length ?? 0} ball(s)
            </span>
            <div class="galton-board__column-fill" style="height: {column?.length ?? 0}px" />
          </div>
          <p>Column {idx + 1}</p>
        </div>
      {/each}
    </div>
  </div>
{/if}

<script lang="ts">
  import type { GaltonBall } from './types/GaltonBall';

  import type { GaltonBoard } from './types/GaltonBoard';
  import { SocketClientMessage, SocketServerMessage, ClientMessages } from './utilities/sockets';

  let socket: WebSocket;
  let socketOpen: number;
  let reconnectInterval: NodeJS.Timer;
  let board: GaltonBoard;
  let ballCount = 0;

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

      const message: SocketClientMessage = {
        msg: ClientMessages.Opened,
      };

      socket.send(JSON.stringify(message));

      socketOpen = socket.readyState;
    });

    socket.addEventListener('message', (ev) => {
      const parsedData: SocketServerMessage = JSON.parse(ev.data);

      if (parsedData.MessageName === ClientMessages.BoardState) {
        board = parsedData.MessageData;

        if (board.Columns) {
          ballCount = board.Columns.reduce((acc: number, curr: Array<GaltonBall>) => {
            return curr ? acc + curr.length : acc;
          }, 0);
        }
      }
    });

    socket.addEventListener('close', (_) => {
      pollConnection();

      socketOpen = socket.CLOSED;
    });
  };

  const handleFormSubmit = (e: SubmitEvent) => {
    const submittedFormElm = <HTMLFormElement>e.currentTarget;
    const boardDepth = submittedFormElm.elements['quantity'].value;

    const message: SocketClientMessage = {
      msg: ClientMessages.CreateBoard,
      quantity: Number.parseInt(boardDepth),
    };

    socket.send(JSON.stringify(message));
  };

  const handleAddBalls = (e: SubmitEvent) => {
    const submittedFormElm = <HTMLFormElement>e.currentTarget;
    const quantityToAdd = submittedFormElm.elements['quantity'].value;

    const message: SocketClientMessage = {
      msg: ClientMessages.AddBalls,
      quantity: Number.parseInt(quantityToAdd),
    };

    socket.send(JSON.stringify(message));
  };

  const handleResetBoard = () => {
    const message: SocketClientMessage = {
      msg: ClientMessages.ResetBoard,
      quantity: 0,
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
    margin-bottom: 0;
  }

  .board-create-form {
    label {
      color: #aaaaaa;
    }
  }

  .galton-board {
    align-items: flex-end;
    display: flex;
    height: 800px;
    justify-content: space-between;
    overflow: hidden;

    span {
      font-size: 0.75em;
      z-index: 1;
    }

    &__wrap {
      background-color: #565656;
      border: 3px solid black;
      border-radius: 20px;
      margin: 1em auto 2em;
      max-width: 1500px;
      padding: 1em 2em;

      h2 {
        font-size: 2em;
        margin: 0 auto 0.5em;
      }

      form {
        margin: 1em auto;
      }
    }

    &__reset {
      margin: 0em auto 1em;
    }

    &__column-wrap {
      display: flex;
      flex-direction: column;
      height: 100%;
    }

    &__column {
      align-items: center;
      border: 3px solid black;
      border-radius: 20px;
      display: flex;
      justify-content: center;
      flex: 1;
      overflow: hidden;
      position: relative;
    }

    &__column-fill {
      background-color: #ff3e00;
      bottom: 0;
      position: absolute;
      left: 0;
      right: 0;
    }
  }
</style>
