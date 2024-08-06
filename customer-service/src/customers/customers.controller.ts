import { Body, Controller, Get, HttpStatus, Param, Patch, Query, Request, UseGuards } from '@nestjs/common';
import { CustomersService } from './customers.service';
import { JwtAuthGuard } from '../auth/guards/jwt-auth.guard';
import { ResponseObject } from '../response/response-object.dto';
import { RedisCacheService } from '../redis/redis.service';
import { allCustomerKey, customerKey } from '../redis/key';
import { UpdateAccountDto } from './dto/update-account.dto';
import { ResponsePaginationDTO } from './dto/response-pagination.dto';
import { Customer } from './models/customer.schema';
import { RequestPaginationDto } from './dto/request-pagination.dto';

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

  @UseGuards(JwtAuthGuard)
  @Get()
  async getAllCustomers(@Query() pagination: RequestPaginationDto): Promise<ResponseObject> {
    const { keyword, page, limit, sort } = pagination;

    if (!keyword) {
      const cachedAllUsers = await this.redisCacheService.get(
        allCustomerKey(page, limit, sort),
      );
      if (cachedAllUsers) {
        return {
          data: cachedAllUsers,
          status: HttpStatus.OK,
          message: 'Get all users successfully',
        };
      }
    }

    const customers = await this.customersService.getAllCustomers(pagination)
    const paginationDto = this.buildPaginationDto(pagination, customers);

    customers.length !== 0 && this.redisCacheService.set(allCustomerKey(page, limit, sort), paginationDto);

    return {
      data: paginationDto,
      status: HttpStatus.OK,
      message: 'Get all customers successfully',
    };
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  async getUser(@Param('id') _id: string): Promise<ResponseObject> {
    const customer = await this.customersService.findByCustomerId(_id);

    return {
      data: customer,
      status: HttpStatus.OK,
      message: 'Get customer successfully',
    };
  }

  private buildPaginationDto(
    pagination: RequestPaginationDto,
    customers: Customer[],
  ): ResponsePaginationDTO {
    const { page, limit } = pagination;

    return {
      total_items: customers.length,
      total_pages: Math.ceil(customers.length / limit),
      current_page: parseInt(String(page)),
      start_count: (page - 1) * limit + 1,
      end_count: page * limit > customers.length ? customers.length : page * limit,
      entities: customers,
    };
  }

}
