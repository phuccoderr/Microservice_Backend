import { IsNotEmpty, IsString } from 'class-validator';

class Person {
  id: string;
  name: string;
  email: string;
}

export class UserSendMessageDto {
  sender: Person;

  receiver: Person;

  @IsString()
  @IsNotEmpty()
  message: string;
}
