import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { AuthController } from './auth.controller';
import { DatabaseModule } from '../database/database.module';
import { CustomerRefreshToken, CustomerRefreshTokenSchema } from './models/refresh-token.schema';
import { JwtStrategy } from './strategies/jwt.strategy';
import { CustomersModule } from '../customers/customers.module';

@Module({
  imports: [
    DatabaseModule,
    DatabaseModule.forFeature([{name: CustomerRefreshToken.name, schema: CustomerRefreshTokenSchema}]),
    CustomersModule
  ],
  controllers: [AuthController],
  providers: [AuthService, JwtStrategy],
})
export class AuthModule {}
