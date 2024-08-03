import { Body, Controller, Get, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateCustomerDto } from './dto/create-customer.dto';
import { ResponseObject } from '../response/response-object.dto';
import { ProducerService } from '../kafka/producer.service';
import { VerifyPayLoad } from '../kafka/dto/kafka-verify.dto';

@Controller('api/v1/customers')
export class AuthController {
  constructor(private readonly authService: AuthService,
              private readonly producerService: ProducerService) {}

  @Post('register')
  async register(@Body() createCustomerDto: CreateCustomerDto, ): Promise<ResponseObject> {

    const user = await this.authService.register(createCustomerDto);

    const payload: VerifyPayLoad = {
      email: user.email,
      url: "https://localhost:9150/api/v1/customers/verify?token" + user.verification_code
    }

   this.producerService.produce("customer-verify-events-topics", {
     value: JSON.stringify(payload),
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
