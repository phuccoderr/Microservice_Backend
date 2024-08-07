import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';
import { ROLE } from "@src/auth/decorators/role.enum";



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

  @Prop({type: [String],enum: ROLE, default: ROLE.USER})
  roles?: ROLE[];
}

export const UserSchema = SchemaFactory.createForClass(User);
