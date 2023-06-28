import { Injectable } from '@angular/core';
import * as socketIo from 'socket.io-client';

const SERVER_URL = 'http://localhost:3000';

@Injectable({
  providedIn: 'root'
})
export class SocketServiceService {

  private socket: any;
  constructor() {
    this.socket = socketIo(SERVER_URL);
   }
}
  +
