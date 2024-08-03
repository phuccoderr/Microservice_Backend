import { Injectable, NotFoundException } from '@nestjs/common';
import { Customer } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';
import { UpdateAccountDto } from './dto/update-account.dto';

@Injectable()
export class CustomersService {

  constructor(private readonly customersRepository: CustomersRepository) {}

  async findByCustomerId(customerId: string): Promise<Customer> {
      const customer = await this.customersRepository.findOne({_id: customerId}, "-password");
      if (!customer) {
        throw new NotFoundException('Customer not found');
      }

      return customer
  }

  async updateCustomer(customerId: string, updateAccountDto: UpdateAccountDto): Promise<Customer> {
    const customer = await this.customersRepository.findOneAndUpdate({_id: customerId}, updateAccountDto);
    if (!customer) {
      throw new NotFoundException('Customer not found');
    }

    return customer
  }
}
