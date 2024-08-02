import { Injectable, NotFoundException } from '@nestjs/common';
import { Customer } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';

@Injectable()
export class CustomersService {

  constructor(private readonly customersRepository: CustomersRepository) {}

  async findByCustomerId(customerId: string): Promise<Customer> {
    try {
      return await this.customersRepository.findOne({_id: customerId}, "-password");
    } catch (error) {
      throw new NotFoundException('Customer not found');
    }
  }
}
