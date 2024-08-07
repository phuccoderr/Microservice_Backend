import { IsArray, IsNotEmpty, IsOptional, IsString } from 'class-validator';
import { ROLE } from "@src/auth/decorators/role.enum";

export class UpdateUserDto {
  @IsString()
  @IsNotEmpty()
  name: string;

  @IsOptional()
  @IsArray()
  @IsString({ each: true })
  @IsNotEmpty({ each: true })
  roles?: ROLE[];
}
