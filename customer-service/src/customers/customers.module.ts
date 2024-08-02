import { Module } from '@nestjs/common';
import { CustomersService } from './customers.service';
import { CustomersController } from './customers.controller';
import { DatabaseModule } from '../database/database.module';
import { Customer, CustomerSchema } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';

@Module({
  imports: [
    DatabaseModule,
    DatabaseModule.forFeature([{name: Customer.name, schema: CustomerSchema}]),
  ],
  controllers: [CustomersController],
  providers: [CustomersService, CustomersRepository],
  exports: [CustomersRepository]
})
export class CustomersModule {}
