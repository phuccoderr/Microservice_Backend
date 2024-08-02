import { Body, Controller, Get, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateCustomerDto } from './dto/create-customer.dto';
import { ResponseObject } from '../response/response-object.dto';
import { KafkaService } from '../kafka/kafka.service';

@Controller('api/v1/customers')
export class AuthController {
  constructor(private readonly authService: AuthService,
              private readonly producerService: KafkaService) {}

  @Post('register')
  async register(@Body() createCustomerDto: CreateCustomerDto, ): Promise<ResponseObject> {

    const user = await this.authService.register(createCustomerDto);

    this.producerService.sendMessage('customer-verify-events-topics', {
      value: 'Hello World',
    })

    return {
      data: user,
      status: HttpStatus.CREATED,
      message: 'register successfully',
    }
  }

  @Get('verify')
  async verify(@Query('token') token: string): Promise<ResponseObject> {

    await this.authService.verify(token);

    return {
      data: {},
      status: HttpStatus.OK,
      message: 'verification successfull',
    }
  }
}
