import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';

export enum AuthenticationType {
  DATABASE= "DATABASE",
  FACEBOOK = "FACEBOOK",
  GOOGLE = "GOOGLE"
}

@Schema({versionKey: false})
export class Customer extends AbstractDocument {
  @Prop({ name: 'email', unique: true, required: true })
  email: string;

  @Prop({name: 'password'})
  password: string;

  @Prop({name: 'status', default: false})
  status: boolean;

  @Prop({name: 'verification_code'})
  verification_code: string;

  @Prop({name: 'reset_password_token'})
  reset_password_token: string;

  @Prop({name: 'authentication_type', type: String, enum: AuthenticationType, default: AuthenticationType.DATABASE})
  authentication_type: AuthenticationType;

  @Prop({name: 'first_name', required: true})
  first_name: string;

  @Prop({name: 'last_name', required: true})
  last_name: string;

  @Prop({name: 'phone_number'})
  phone_number: string;

  @Prop({name: 'address'})
  address: string;
}

export const CustomerSchema = SchemaFactory.createForClass(Customer);