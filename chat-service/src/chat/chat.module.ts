import { Module } from '@nestjs/common';
import { ChatController } from 'src/chat/chat.controller';
import { ChatService } from 'src/chat/chat.service';

@Module({
  controllers: [ChatController],
  providers: [ChatService],
})
export class ChatModule {}
