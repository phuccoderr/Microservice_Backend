import { Injectable, NotFoundException, UnprocessableEntityException } from '@nestjs/common';
import { CreateCustomerDto } from './dto/create-customer.dto';
import * as bcrypt from 'bcrypt';
import { AuthenticationType, Customer } from '../customers/models/customer.schema';
import { CustomersRepository } from '../customers/customers.repository';
import { v4 as uuidv4 } from 'uuid';

@Injectable()
export class AuthService {

  constructor(private readonly customersRepository: CustomersRepository) {
  }

  async register(createCustomerDto: CreateCustomerDto): Promise<Customer> {
    try {
      return await this.customersRepository.create({
        address: '',
        authentication_type: AuthenticationType.DATABASE,
        phone_number: '',
        reset_password_token: '',
        status: false,
        verification_code: uuidv4(),
        ...createCustomerDto,
        password: await bcrypt.hash(createCustomerDto.password, 10)
      })
    } catch (error) {
      throw new UnprocessableEntityException(error.message);
    }
  }

  async verify (code: string): Promise<void> {
    console.log(code)
    const customer = await this.customersRepository.findByVerificationCode(code);
    if (!customer || customer.status) {
      throw new NotFoundException('Verify fail');
    }

    await this.customersRepository.findOneAndUpdate(
      {_id: customer._id},
      {verification_code: '', status: true});
  }
}
