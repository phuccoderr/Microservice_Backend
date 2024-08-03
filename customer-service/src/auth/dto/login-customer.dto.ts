import { IsEmail, IsStrongPassword } from 'class-validator';

export class LoginCustomerDto {
  @IsEmail()
  email: string;

  @IsStrongPassword()
  password: string;
}