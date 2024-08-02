import { IsEmail, IsNotEmpty, IsString, IsStrongPassword } from 'class-validator';
import { Expose } from 'class-transformer';

export class CreateCustomerDto {
  @IsNotEmpty()
  @IsEmail()
  email: string;

  @IsStrongPassword()
  password: string;

  @IsString()
  @IsNotEmpty()
  last_name: string;

  @IsString()
  @IsNotEmpty()
  first_name: string;
}