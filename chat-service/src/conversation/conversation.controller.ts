import { HttpService } from '@nestjs/axios';
import { Controller, Get, HttpStatus, Param, UseGuards } from '@nestjs/common';
import { JwtAuthGuard } from '@src/auth/guards/jwt-auth.guard';
import { CHAT_CONSTANTS } from '@src/constants/chat-constants';
import { ConversationService } from '@src/conversation/conversation.service';
import { ResponseObject } from '@src/response/response-object.dto';

@Controller('api/v1/chats/conversations')
export class ConversationController {
  constructor(private readonly conversationService: ConversationService) {}

  // @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getOne(@Param('id') id: string): Promise<ResponseObject> {
    const conversations = await this.conversationService.getAll(id);

    return {
      data: conversations,
      status: HttpStatus.OK,
      message: CHAT_CONSTANTS.GET_CONVERSATION,
    };
  }
}
