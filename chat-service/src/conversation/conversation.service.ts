import { SERVICE_CONSTANTS } from './../constants/service_constants';
import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { ConversationRepository } from '@src/conversation/conversation.repository';
import { ConversationDto } from '@src/conversation/dto/conversation.dto';
import { UserSendMessageDto } from '@src/events/dto/user-message.dto';

@Injectable()
export class ConversationService {
  constructor(
    private readonly repository: ConversationRepository,
    private readonly httpService: HttpService,
  ) {}

  async getConversation({ sender_id, receiver_id }: ConversationDto) {
    return await this.repository.get({ sender_id, receiver_id });
  }

  async getAll(user_id: string) {
    console.log(user_id);
    return await this.repository.findAll(user_id);
  }

  async createConversation(data: UserSendMessageDto) {
    return await this.repository.create(data);
  }
}
