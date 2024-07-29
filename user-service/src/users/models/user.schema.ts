import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';

@Schema({ versionKey: false })
export class User extends AbstractDocument {
  @Prop({ unique: true })
  email: string;

  @Prop({ required: true })
  password: string;

  @Prop({ required: true })
  name: string;

  @Prop({ required: true, default: true })
  status: boolean;

  @Prop()
  roles?: string[];
}

export const UserSchema = SchemaFactory.createForClass(User);
