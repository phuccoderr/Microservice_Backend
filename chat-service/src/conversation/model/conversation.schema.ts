import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Message } from '@src/chat/model/message.schema';
import { AbstractDocument } from '@src/database/abstract.schema';
import { SchemaTypes, Types } from 'mongoose';

@Schema()
export class Conversation extends AbstractDocument {
  @Prop({
    ref: 'Customer',
  })
  participants: Types.ObjectId[];

  @Prop([{ type: SchemaTypes.ObjectId, ref: Message.name }])
  messages: Types.ObjectId[];

  @Prop({})
  created_at: Date;
}

export const ConversationSchema = SchemaFactory.createForClass(Conversation);
