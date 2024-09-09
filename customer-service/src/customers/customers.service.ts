import {
  BadRequestException,
  HttpStatus,
  Injectable,
  Logger,
  NotFoundException,
} from '@nestjs/common';
import { Customer } from './models/customer.schema';
import { CustomersRepository } from './customers.repository';
import { UpdateAccountDto } from './dto/update-account.dto';
import { RequestPaginationDto } from './dto/request-pagination.dto';
import { DATABASE_CONST } from '../constants/db-constants';
import { CUSTOMER_CONSTANTS } from '@src/constants/customer-constants';
import { ChangePasswordDto } from '@src/customers/dto/change-password.dto';
import * as bcrypt from 'bcrypt';
import { AUTH_CONSTANTS } from '@src/constants/auth-constants';

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

  async updateStatus(_id: string, status: string) {
    if(status != 'true' && status != 'false') {
      this.logger.warn('Status is true or false!');
      throw new BadRequestException(CUSTOMER_CONSTANTS.STATUS);
    }

    try {
      await this.customersRepository.findOneAndUpdate({_id}, {status});
    } catch (error) {
      this.logger.warn('Customer does not exist with id ' + _id);
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }
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

  async changePassword(_id: string, changePassword: ChangePasswordDto): Promise<void> {
    const customer = await this.customersRepository.findById(_id);
    if (!customer) {
      this.logger.warn('customer not found!');
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }

    const isMatch = await bcrypt.compare(changePassword.old_password, customer.password);
    if (!isMatch) {
      this.logger.log('password not match!');
      throw new BadRequestException(AUTH_CONSTANTS.PASSWORD_NOT_MATCH);
    }

    await this.customersRepository.findOneAndUpdate(
      {_id},
      {password: await bcrypt.hash(changePassword.new_password, 10)}
    );

  }
}
