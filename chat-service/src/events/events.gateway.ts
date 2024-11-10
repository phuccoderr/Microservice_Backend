import { UserSendMessageDto } from './dto/user-message.dto';
import {
  MessageBody,
  OnGatewayConnection,
  OnGatewayDisconnect,
  SubscribeMessage,
  WebSocketGateway,
  WebSocketServer,
} from '@nestjs/websockets';
import { MessageService } from '@src/message/message.service';
import { Server, Socket } from 'socket.io';

@WebSocketGateway({
  cors: {
    origin: ['http://localhost:4000'],
  },
})
export class EventsGateway implements OnGatewayConnection, OnGatewayDisconnect {
  constructor(private readonly messageService: MessageService) {}

  @WebSocketServer() server: Server;

  handleDisconnect(socket: Socket) {
    console.log('Client disconnected:', socket.id);
  }
  handleConnection(socket: Socket) {
    console.log('Client connected:', socket.id);
  }

  @SubscribeMessage('send-messages')
  async handleMessage(@MessageBody() data: UserSendMessageDto) {
    const response = await this.messageService.sendMessage(data);

    this.server.emit('receive-messages', response);
  }

  @SubscribeMessage('typing')
  async handleTyping(@MessageBody() id: string) {
    this.server.emit('typing', id);
  }

  @SubscribeMessage('notyping')
  async handleNoTyping(@MessageBody() id: string) {
    this.server.emit('notyping', id);
  }
}
