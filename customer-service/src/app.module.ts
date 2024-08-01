import { Module } from '@nestjs/common';
import { CustomersModule } from './customers/customers.module';
import { ConfigModule } from "@nestjs/config";
import { LoggerModule as PinoLoggerModule } from "nestjs-pino";
import { DatabaseModule } from './database/database.module';
import { AuthModule } from './auth/auth.module';

@Module({
  imports: [
    DatabaseModule,
    CustomersModule,
    ConfigModule.forRoot({
      isGlobal: true
    }),
    PinoLoggerModule.forRoot({
      pinoHttp: {
        transport: {
          target: 'pino-pretty',
          options: {
            singleLine: true
          }
        }
      }
    }),
    DatabaseModule,
    AuthModule
  ],
})
export class AppModule {}
