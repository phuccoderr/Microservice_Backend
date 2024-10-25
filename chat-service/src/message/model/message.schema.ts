import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '@src/database/abstract.schema';
import { SchemaTypes, Types } from 'mongoose';

class Person {
  id: string;
  name: string;
  email: string;
}

@Schema({ versionKey: false })
export class Message extends AbstractDocument {
  @Prop({ type: Person, required: true })
  sender: Person;

  @Prop({ type: Person, required: true })
  receiver: Person;

  @Prop({ type: String, required: true })
  message: string;

  @Prop({
    type: Date,
    default: new Date(),
  })
  created_at: Date;
}

export const MessageSchema = SchemaFactory.createForClass(Message);
