import { Injectable, NotFoundException } from '@nestjs/common';
import { Customer } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';
import { UpdateAccountDto } from './dto/update-account.dto';
import { RequestPaginationDto } from './dto/request-pagination.dto';

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

  async updateCustomer(customerId: string, updateAccountDto: UpdateAccountDto): Promise<Customer> {
    const customer = await this.customersRepository.findOneAndUpdate({_id: customerId}, updateAccountDto);
    if (!customer) {
      throw new NotFoundException('Customer not found');
    }

    return customer
  }

  async getAllCustomers(query: RequestPaginationDto): Promise<Customer[]> {
    const { keyword, page, limit, sort } = query;

    if (keyword) {
      const filter = { email: keyword, name: keyword };
      return await this.customersRepository.search(page, limit, sort, filter);
    }

    return await this.customersRepository.listByPage(page, limit, sort);
  }
}
