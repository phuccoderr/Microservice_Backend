import { Injectable } from '@nestjs/common';
import { ConversationService } from '@src/conversation/conversation.service';
import { UserSendMessageDto } from '@src/events/dto/user-message.dto';
import { MessageRepository } from '@src/message/message.repository';
import { Types } from 'mongoose';

@Injectable()
export class MessageService {
  constructor(
    private readonly repository: MessageRepository,
    private readonly conversationService: ConversationService,
  ) {}

  async sendMessage(sendMessage: UserSendMessageDto) {
    const { sender, receiver, message } = sendMessage;

    let conversation = await this.conversationService.getConversation({
      sender_id: sender.id,
      receiver_id: receiver.id,
    });

    console.log('conversation', conversation);

    if (!conversation) {
      conversation =
        await this.conversationService.createConversation(sendMessage);
    }

    const newMessage = await this.repository.create({
      _id: new Types.ObjectId(),
      sender: {
        id: sender.id,
        name: sender.name,
        email: sender.email,
      },
      receiver: {
        id: receiver.id,
        name: receiver.name,
        email: receiver.email,
      },
      message,
      created_at: new Date(),
    });

    if (newMessage) {
      conversation.messages.push(newMessage._id);
    }

    await conversation.save();
    return await newMessage.save();
  }

  async getMessages(sender_id: string, receiver_id: string) {
    let conversation = await this.conversationService.getConversation({
      sender_id,
      receiver_id,
    });

    return conversation;
  }
}
