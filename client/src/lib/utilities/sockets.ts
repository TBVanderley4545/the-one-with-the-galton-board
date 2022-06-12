import type { GaltonBoard } from '../types/GaltonBoard';

export enum ClientMessages {
  Opened = 'new observer',
  CreateBoard = 'create board',
  BoardState = 'board state',
  AddBalls = 'add balls',
}

export interface SocketClientMessage {
  msg: string;
  quantity?: number;
}

export interface SocketServerMessage {
  MessageName: string;
  MessageData: GaltonBoard;
}
