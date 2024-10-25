import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AuthModule } from '@src/auth/auth.module';
import { ConversationModule } from '@src/conversation/conversation.module';
import { DatabaseModule } from '@src/database/database.module';
import { EventsModule } from '@src/events/events.module';
import { MessageModule } from '@src/message/message.module';
import { LoggerModule as PinoLoggerModule } from 'nestjs-pino';

@Module({
  imports: [
    DatabaseModule,
    MessageModule,
    ConversationModule,
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
    HttpModule,
  ],
})
export class AppModule {}
