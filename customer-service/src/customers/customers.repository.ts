import { Injectable } from '@nestjs/common';
import { Logger } from '@nestjs/common';
import { AbstractRepository } from '../database/abstract.repository';
import { Customer } from './models/customer.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

@Injectable()
export class CustomersRepository extends AbstractRepository<Customer> {
  protected readonly logger: Logger = new Logger(Customer.name);

  constructor(@InjectModel(Customer.name) customerModel: Model<Customer>) {
    super(customerModel);
  }

  async findByVerificationCode(code: string): Promise<Customer> {
    return this.model.findOne({ verification_code: code }).lean<Customer>(true);
  }
}