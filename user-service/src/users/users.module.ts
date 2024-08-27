import { UsersRepository } from './repository/users.repository';
import { Module } from '@nestjs/common';
import { UsersService } from './users.service';
import { UsersController } from './users.controller';
import { DatabaseModule } from '../database/database.module';
import { User, UserSchema } from './models/user.schema';
import { RedisModule } from '../redis/redis.module';

@Module({
  imports: [
    DatabaseModule,
    RedisModule,
    DatabaseModule.forFeature([{ name: User.name, schema: UserSchema }]),
  ],
  controllers: [UsersController],
  providers: [UsersService, UsersRepository],
  exports: [UsersService, UsersRepository],
})
export class UsersModule {}
