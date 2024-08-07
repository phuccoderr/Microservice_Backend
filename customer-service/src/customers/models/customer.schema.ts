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
      'https://res-console.cloudinary.com/dp4tp9gwa/thumbnails/v1/image/upload/v1722999166/YXZhdGFyLWRlZmF1bHRfYWg5eHln/drilldown',
  })
  avatar?: string;

  @Prop({ unique: true })
  image_id?: string;

  @Prop({ default: false })
  status: boolean;

  @Prop({ unique: true })
  verification_code: string;

  @Prop({ unique: true })
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
}

export const CustomerSchema = SchemaFactory.createForClass(Customer);
