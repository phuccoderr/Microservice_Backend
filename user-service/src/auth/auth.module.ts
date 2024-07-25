import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { AuthController } from './auth.controller';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { UsersModule } from 'src/users/users.module';
import { JwtStrategy } from 'src/auth/strategies/jwt.strategy';
import { DatabaseModule } from 'src/database/database.module';
import {
  RefreshToken,
  RefreshTokenSchema,
} from 'src/auth/models/refresh-token.schema';
import { RefreshTokenRepository } from 'src/auth/refresh-token.repository';

@Module({
  imports: [
    UsersModule,
    DatabaseModule.forFeature([
      { name: RefreshToken.name, schema: RefreshTokenSchema },
    ]),
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    JwtModule.registerAsync({
      useFactory: (configService: ConfigService) => ({
        secret: configService.get('JWT_SECRET'),
        signOptions: {
          expiresIn: `${configService.get('JWT_EXPIRATION')}s`,
        },
      }),
      inject: [ConfigService],
    }),
  ],
  controllers: [AuthController],
  providers: [AuthService, JwtStrategy, RefreshTokenRepository],
})
export class AuthModule {}
