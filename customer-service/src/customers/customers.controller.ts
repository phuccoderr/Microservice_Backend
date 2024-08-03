import { Body, Controller, Get, HttpStatus, Patch, Request, UseGuards } from '@nestjs/common';
import { CustomersService } from './customers.service';
import { JwtAuthGuard } from '../auth/guards/jwt-auth.guard';
import { ResponseObject } from '../response/response-object.dto';
import { RedisCacheService } from '../redis/redis.service';
import { customerKey } from '../redis/key';
import { UpdateAccountDto } from './dto/update-account.dto';

@Controller('/api/v1/customers')
export class CustomersController {
  constructor(private readonly customersService: CustomersService,
              private readonly redisCacheService: RedisCacheService,) {}

  @UseGuards(JwtAuthGuard)
  @Get('account')
  async getAccount(@Request() req): Promise<ResponseObject> {
    const { _id } = req.user;

    const cacheCustomer = await this.redisCacheService.get(customerKey(_id));
    if (cacheCustomer) {
      return {
        data: cacheCustomer,
        status: HttpStatus.OK,
        message: 'Get Account Successfully!',
      };
    }
    const customer = await this.customersService.findByCustomerId(_id);
    this.redisCacheService.set(customerKey(_id), customer);

    return {
      data: customer,
      status: HttpStatus.OK,
      message: 'Get Account Successfully!',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Patch('account')
  async updateAccount(@Request() req,@Body() updateAccountDto: UpdateAccountDto): Promise<ResponseObject> {
    const { _id } = req.user;
    const customer = await this.customersService.updateCustomer(_id,updateAccountDto);
    this.redisCacheService.del(customerKey(_id));

    this.redisCacheService.set(_id, customer);

    return {
      data: customer,
      status: HttpStatus.OK,
      message: 'Get Account Successfully!',
    };
  }
}
