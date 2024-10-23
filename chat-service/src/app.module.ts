import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AuthModule } from '@src/auth/auth.module';
import { ChatModule } from '@src/chat/chat.module';
import { DatabaseModule } from '@src/database/database.module';
import { EventsModule } from '@src/events/events.module';
import { LoggerModule as PinoLoggerModule } from 'nestjs-pino';

@Module({
  imports: [
    ChatModule,
    DatabaseModule,
    AuthModule,
    EventsModule,
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    PinoLoggerModule.forRoot({
      pinoHttp: {
        transport: {
          target: 'pino-pretty',
          options: {
            singleLine: true,
          },
        },
      },
    }),
  ],
})
export class AppModule {}
