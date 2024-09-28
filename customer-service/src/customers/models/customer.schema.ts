import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';

export enum AuthenticationType {
  DATABASE = 'DATABASE',
  FACEBOOK = 'FACEBOOK',
  GOOGLE = 'GOOGLE',
}

@Schema({ versionKey: false })
export class Customer extends AbstractDocument {
  @Prop({ unique: true, required: true })
  email: string;

  @Prop()
  password: string;

  @Prop({
    default:
      '',
  })
  avatar?: string;

  @Prop()
  image_id?: string;

  @Prop({ default: false })
  status: boolean;

  @Prop()
  verification_code: string;

  @Prop()
  reset_password_token?: string;

  @Prop({
    type: String,
    enum: AuthenticationType,
    default: AuthenticationType.DATABASE,
  })
  authentication_type: AuthenticationType;

  @Prop({ required: true })
  first_name: string;

  @Prop({ required: true })
  last_name: string;

  @Prop()
  phone_number?: string;

  @Prop()
  address?: string;

  @Prop({type: [String], default: "CUSTOMER"})
  roles?: string[];
}

export const CustomerSchema = SchemaFactory.createForClass(Customer);
