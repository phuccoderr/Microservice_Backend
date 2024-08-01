import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { AuthController } from './auth.controller';
import { DatabaseModule } from '../database/database.module';
import { CustomerRefreshToken, CustomerRefreshTokenSchema } from './models/refresh-token.schema';

@Module({
  imports: [
    DatabaseModule,
    DatabaseModule.forFeature([{name: CustomerRefreshToken.name, schema: CustomerRefreshTokenSchema}]),
  ],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
