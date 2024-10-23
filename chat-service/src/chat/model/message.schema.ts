import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '@src/database/abstract.schema';
import { SchemaTypes, Types } from 'mongoose';

@Schema()
export class Message extends AbstractDocument {
  @Prop({ type: SchemaTypes.ObjectId, required: true })
  sender_id: Types.ObjectId;

  @Prop({ type: SchemaTypes.ObjectId, required: true })
  receiver_id: Types.ObjectId;

  @Prop({ type: String, required: true })
  message: string;

  @Prop({
    type: Date,
    default: new Date(),
  })
  created_at: Date;
}

export const MessageSchema = SchemaFactory.createForClass(Message);
