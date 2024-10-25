import { Module } from '@nestjs/common';
import { EventsGateway } from '@src/events/events.gateway';
import { MessageModule } from '@src/message/message.module';

@Module({
  imports: [MessageModule],
  providers: [EventsGateway],
})
export class EventsModule {}
