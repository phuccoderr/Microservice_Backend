import { IsStrongPassword } from 'class-validator';

export class ChangePasswordDto {
  @IsStrongPassword()
  old_password: string;

  @IsStrongPassword()
  new_password: string;
}