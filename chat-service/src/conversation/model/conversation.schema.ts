import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '@src/database/abstract.schema';
import { Message } from '@src/message/model/message.schema';
import { SchemaTypes, Types } from 'mongoose';

class Person {
  id: string;
  name: string;
  email: string;
}

@Schema({ versionKey: false })
export class Conversation extends AbstractDocument {
  @Prop({ type: [Person] })
  participants: Person[];

  @Prop([{ type: SchemaTypes.ObjectId, ref: Message.name, default: [] }])
  messages?: Types.ObjectId[];

  @Prop({
    type: Date,
    default: new Date(),
  })
  created_at: Date;
}

export const ConversationSchema = SchemaFactory.createForClass(Conversation);
