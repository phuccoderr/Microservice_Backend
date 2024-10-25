import {
  Controller,
  Get,
  HttpStatus,
  Param,
  Query,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from '@src/auth/guards/jwt-auth.guard';
import { CHAT_CONSTANTS } from '@src/constants/chat-constants';
import { MessageService } from '@src/message/message.service';
import { ResponseObject } from '@src/response/response-object.dto';

@Controller('api/v1/chats/messages')
export class MessageController {
  constructor(private readonly messageService: MessageService) {}

  // @UseGuards(JwtAuthGuard)
  @Get('/:id')
  async getMessages(
    @Param('id') id: string,
    @Query('receiver_id') receiver_id: string,
  ): Promise<ResponseObject> {
    const messages = await this.messageService.getMessages(id, receiver_id);

    return {
      data: messages,
      status: HttpStatus.OK,
      message: CHAT_CONSTANTS.GET_MESSAGES,
    };
  }
}
