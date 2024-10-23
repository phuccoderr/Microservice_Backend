import {
  OnGatewayConnection,
  OnGatewayDisconnect,
  WebSocketGateway,
  WebSocketServer,
} from '@nestjs/websockets';
import { Server, Socket } from 'socket.io';

@WebSocketGateway({
  cors: {
    origin: ['http://localhost:4000'],
  },
})
export class EventsGateway implements OnGatewayConnection, OnGatewayDisconnect {
  @WebSocketServer() server: Server;

  handleDisconnect(socket: Socket) {
    console.log('Client disconnected:', socket.id);
  }
  handleConnection(socket: Socket) {
    console.log('Client connected:', socket.id);
  }
}
