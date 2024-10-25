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
        participants: { $all: [sender_id, receiver_id] },
      })
      .populate('messages');
  }

  async findAll(user_id: string) {
    return await this.model
      .find({
        participants: { $in: [user_id] },
      })
      .populate('messages');
  }

  async create({ sender, receiver }: UserSendMessageDto) {
    return await this.model.create({
      participants: [sender.id, receiver.id],
      _id: new Types.ObjectId(),
    });
  }
}
