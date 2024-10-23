import { Module } from '@nestjs/common';
import { EventsGateway } from '@src/events/events.gateway';

@Module({
  providers: [EventsGateway],
})
export class EventsModule {}
