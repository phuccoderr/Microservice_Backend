import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';
import { Types } from "mongoose";

@Schema({ versionKey: false })
export class UserRefreshToken extends AbstractDocument {
  @Prop({ required: true, unique: true })
  token: string;

  @Prop({ type: Types.ObjectId, required: true, ref: "User" })
  userId: string;

  @Prop({ required: true })
  expiresAt: Date;
}

export const UserRefreshTokenSchema = SchemaFactory.createForClass(UserRefreshToken);
