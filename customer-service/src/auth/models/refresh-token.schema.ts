import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Types } from 'mongoose';

@Schema({versionKey: false})
export class CustomerRefreshToken {

  @Prop({type: Types.ObjectId})
  _id: Types.ObjectId;

  @Prop({ required: true, unique: true })
  token: string;

  @Prop({type: Types.ObjectId, required: true , ref: 'Customer'})
  customerId: string

  @Prop({ required: true })
  expiresAt: Date;
}

export const CustomerRefreshTokenSchema = SchemaFactory.createForClass(CustomerRefreshToken);