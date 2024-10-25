import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';

import { Message } from '@src/message/model/message.schema';
import { Model } from 'mongoose';

@Injectable()
export class MessageRepository {
  constructor(
    @InjectModel(Message.name) private readonly model: Model<Message>,
  ) {}

  async create(message: Message) {
    return await new this.model(message);
  }
}
