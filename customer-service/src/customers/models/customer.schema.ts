import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '../../database/abstract.schema';

const authenticationType = ["FACEBOOK", "GOOGLE","DATABASE"]

@Schema({versionKey: false})
export class Customer extends AbstractDocument {
  @Prop({ name: 'email', unique: true, required: true })
  email: string;

  @Prop({name: 'password'})
  password: string;

  @Prop({name: 'status', default: false})
  status: boolean;

  @Prop({name: 'verification_code'})
  verificationCode: string;

  @Prop({name: 'reset_password_token'})
  resetPasswordToken: string;

  @Prop({name: 'authentication_type', type: String, enum: authenticationType})
  authenticationType = authenticationType;

  @Prop({name: 'first_name', required: true})
  firstName: string;

  @Prop({name: 'last_name'})
  lastName: string;

  @Prop({name: 'phone_number'})
  phoneNumber: string;

  @Prop({name: 'address'})
  address: string;
}

export const CustomerSchema = SchemaFactory.createForClass(Customer);