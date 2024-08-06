import { Injectable } from '@nestjs/common';
import { Logger } from '@nestjs/common';
import { AbstractRepository } from '../database/abstract.repository';
import { Customer } from './models/customer.schema';
import { InjectModel } from '@nestjs/mongoose';
import { FilterQuery, Model } from 'mongoose';

@Injectable()
export class CustomersRepository extends AbstractRepository<Customer> {
  protected readonly logger: Logger = new Logger(Customer.name);

  constructor(@InjectModel(Customer.name) customerModel: Model<Customer>) {
    super(customerModel);
  }

  async findByVerificationCode(code: string): Promise<Customer> {
    return this.model.findOne({ verification_code: code }).lean<Customer>(true);
  }

  async listByPage(page: number, limit: number, sort: 'asc' | 'desc') {
    return await this.model
      .find()
      .sort({ email: sort })
      .skip((page - 1) * limit)
      .limit(limit)
      .select('-password')
      .lean<Customer[]>(true);
  }

  async search(
    page: number,
    limit: number,
    sort: 'asc' | 'desc',
    filter: FilterQuery<Customer>,
  ) {
    return await this.model
      .find({ $or: [{ first_name: filter.first_name }, {last_name: filter.last_name}, { email: filter.email }] })
      .sort({ name: sort })
      .skip((page - 1) * limit)
      .limit(limit)
      .select('-password')
      .lean<Customer[]>(true);
  }
}