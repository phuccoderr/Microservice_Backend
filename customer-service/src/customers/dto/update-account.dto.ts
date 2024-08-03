import { IsString, MaxLength } from 'class-validator';

export class UpdateAccountDto {
  @IsString()
  first_name: string;

  @IsString()
  last_name: string;

  @IsString()
  @MaxLength(10)
  phone_number: string;

  @IsString()
  address: string;
}