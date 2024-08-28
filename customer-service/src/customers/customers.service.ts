import { Injectable, Logger, NotFoundException } from '@nestjs/common';
import { Customer } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';
import { UpdateAccountDto } from './dto/update-account.dto';
import { RequestPaginationDto } from './dto/request-pagination.dto';
import { DATABASE_CONST } from '../constants/db-constants';

@Injectable()
export class CustomersService {
  private logger = new Logger(CustomersService.name);
  constructor(private readonly customersRepository: CustomersRepository) {}

  async findByCustomerId(customerId: string): Promise<Customer> {
    try {
      return await this.customersRepository.findOne(
        { _id: customerId },
        '-password',
      );
    } catch (error) {
      this.logger.warn('customer not found!');
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }

  async updateCustomer(
    customerId: string,
    updateAccountDto: UpdateAccountDto,
  ): Promise<Customer> {
    const customer = await this.customersRepository.findOneAndUpdate(
      { _id: customerId },
      updateAccountDto,
    );
    if (!customer) {
      this.logger.warn('customer not found!');
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }

    return customer;
  }

  async getAllCustomers(query: RequestPaginationDto): Promise<Customer[]> {
    const { keyword, page, limit, sort } = query;

    if (keyword) {
      const filter = { email: keyword, name: keyword };
      return await this.customersRepository.search(page, limit, sort, filter);
    }

    return await this.customersRepository.listByPage(page, limit, sort);
  }

  async uploadAvatar(
    _id: string,
    avatar: string,
    image_id: string,
  ): Promise<void> {
    try {
      await this.customersRepository.findOneAndUpdate(
        { _id },
        { avatar, image_id },
      );
    } catch (error) {
      this.logger.warn('customer not found!');
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
  }
}
