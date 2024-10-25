import { Module } from '@nestjs/common';
import { ConversationModule } from '@src/conversation/conversation.module';
import { DatabaseModule } from '@src/database/database.module';
import { MessageController } from '@src/message/message.controller';
import { MessageRepository } from '@src/message/message.repository';
import { MessageService } from '@src/message/message.service';
import { Message, MessageSchema } from '@src/message/model/message.schema';
@Module({
  imports: [
    DatabaseModule.forFeature([{ name: Message.name, schema: MessageSchema }]),
    ConversationModule,
  ],
  controllers: [MessageController],
  providers: [MessageService, MessageRepository],
  exports: [MessageService],
})
export class MessageModule {}
