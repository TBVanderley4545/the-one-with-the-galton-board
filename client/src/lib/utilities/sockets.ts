export enum ClientMessages {
  Opened = 'new observer',
  CreateBoard = 'create board',
  BoardState = 'board state',
}

export interface SocketMessage {
  msg: string;
  quantity?: number;
}
