import { Module } from '@nestjs/common';
import { CustomersModule } from './customers/customers.module';
import { ConfigModule } from "@nestjs/config";
import { LoggerModule as PinoLoggerModule } from "nestjs-pino";
import { DatabaseModule } from './database/database.module';
import { AuthModule } from './auth/auth.module';
import { KafkaModule } from './kafka/kafka.module';

@Module({
  imports: [
    DatabaseModule,
    AuthModule,
    CustomersModule,
    KafkaModule,
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
  ],
})
export class AppModule {}
