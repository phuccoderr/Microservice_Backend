import { Injectable, Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { ConversationDto } from '@src/conversation/dto/conversation.dto';
import { Conversation } from '@src/conversation/model/conversation.schema';
import { UserSendMessageDto } from '@src/events/dto/user-message.dto';
import { Model, Types } from 'mongoose';

@Injectable()
export class ConversationRepository {
  protected readonly logger: Logger = new Logger(Conversation.name);

  constructor(
    @InjectModel(Conversation.name) private readonly model: Model<Conversation>,
  ) {}

  async get({ sender_id, receiver_id }: ConversationDto) {
    return await this.model
      .findOne({
        $and: [
          { participants: { $elemMatch: { id: sender_id } } },
          { participants: { $elemMatch: { id: receiver_id } } },
        ],
      })
      .populate('messages');
  }

  async findAll(user_id: string) {
    return await this.model
      .find({
        participants: { $elemMatch: { id: user_id } },
      })
      .populate('messages');
  }

  async create({ sender, receiver }: UserSendMessageDto) {
    return await this.model.create({
      participants: [sender, receiver],
      _id: new Types.ObjectId(),
    });
  }
}
