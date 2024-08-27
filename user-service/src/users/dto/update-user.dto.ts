import { IsEnum, IsNotEmpty, IsOptional, IsString } from 'class-validator';
import { ROLE } from '@src/auth/decorators/role.enum';

export class UpdateUserDto {
  @IsString()
  @IsNotEmpty()
  name: string;

  @IsOptional()
  @IsEnum(ROLE, { each: true })
  roles?: ROLE[];
}
