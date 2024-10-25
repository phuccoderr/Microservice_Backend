import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '@src/database/abstract.schema';
import { Message } from '@src/message/model/message.schema';
import { SchemaTypes, Types } from 'mongoose';

@Schema({ versionKey: false })
export class Conversation extends AbstractDocument {
  @Prop({ type: [SchemaTypes.ObjectId] })
  participants: Types.ObjectId[];

  @Prop([{ type: SchemaTypes.ObjectId, ref: Message.name, default: [] }])
  messages?: Types.ObjectId[];

  @Prop({
    type: Date,
    default: new Date(),
  })
  created_at: Date;
}

export const ConversationSchema = SchemaFactory.createForClass(Conversation);
