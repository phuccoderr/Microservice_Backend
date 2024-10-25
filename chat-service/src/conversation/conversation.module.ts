import { HttpModule, HttpService } from '@nestjs/axios';
import { Module } from '@nestjs/common';
import { ConversationController } from '@src/conversation/conversation.controller';
import { ConversationRepository } from '@src/conversation/conversation.repository';
import { ConversationService } from '@src/conversation/conversation.service';
import {
  Conversation,
  ConversationSchema,
} from '@src/conversation/model/conversation.schema';
import { DatabaseModule } from '@src/database/database.module';

@Module({
  imports: [
    DatabaseModule.forFeature([
      { name: Conversation.name, schema: ConversationSchema },
    ]),
    HttpModule,
  ],
  controllers: [ConversationController],
  providers: [ConversationService, ConversationRepository],
  exports: [ConversationService],
})
export class ConversationModule {}
