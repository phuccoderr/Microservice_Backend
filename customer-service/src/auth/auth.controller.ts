import { Body, Controller, Get, HttpCode, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateCustomerDto } from './dto/create-customer.dto';
import { ResponseObject } from '../response/response-object.dto';
import { ProducerService } from '../kafka/producer.service';
import { VerifyPayLoad } from './dto/kafka-verify.dto';
import { LoginCustomerDto } from './dto/login-customer.dto';
import { JwtPayload } from './dto/jwt-payload.dto';
import { ConfigService } from '@nestjs/config';
import { AUTH_CONSTANTS } from '../constants/auth-constants';

@Controller('api/v1/customers/auth')
export class AuthController {
  constructor(private readonly authService: AuthService,
              private readonly producerService: ProducerService,
              private readonly configService: ConfigService,) {}

  @Post('login')
  @HttpCode(HttpStatus.OK)
  async login(@Body() loginCustomerDto: LoginCustomerDto): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.login(loginCustomerDto)

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.LOGIN,
    };
  }

  @Post('refresh')
  @HttpCode(HttpStatus.OK)
  async refreshToken(@Query('token') token: string): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.refreshToken(token);

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.REFRESH_TOKEN,
    };
  }

  @Post('logout')
  @HttpCode(HttpStatus.OK)
  async logout(@Query('token') token: string): Promise<ResponseObject> {
    await this.authService.logout(token);

    return {
      data: {},
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.LOGOUT,
    };
  }

  @Post('register')
  @HttpCode(HttpStatus.CREATED)
  async register(@Body() createCustomerDto: CreateCustomerDto, ): Promise<ResponseObject> {

    const user = await this.authService.register(createCustomerDto);

    const url = this.configService.get("URI_VERIFY_REGISTER");

    const payload: VerifyPayLoad = {
      email: user.email,
      url: `${url}?token=${user.verification_code}`,
    }

    const kafkaTopicVerify = this.configService.get("KAFKA_TOPIC_VERIFY_ACCOUNT");
    this.producerService.produce(kafkaTopicVerify, {
     value: JSON.stringify(payload),
    })

    return {
      data: url,
      status: HttpStatus.CREATED,
      message: AUTH_CONSTANTS.REGISTER,
    }
  }

  @Get('verify')
  async verify(@Query('token') token: string): Promise<ResponseObject> {

    await this.authService.verify(token);

    return {
      data: {},
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.VERIFY_SUCCESS,
    }
  }

  @Post('forgot_password')
  @HttpCode(HttpStatus.OK)
  async forgotPassword(@Query('email') email: string): Promise<ResponseObject> {
    const customer = await this.authService.forgotPassword(email);

    const url = this.configService.get("URI_VERIFY_PASSWORD");

    const payload: VerifyPayLoad = {
      email: email,
      url: `${url}?token=${customer.reset_password_token}`,
    }
    const kafkaTopicPassword = this.configService.get("KAFKA_TOPIC_VERIFY_PASSWORD");
    this.producerService.produce(kafkaTopicPassword, {
      value: JSON.stringify(payload),
    })

    return {
      data: url,
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.FORGOT_PASSWORD,
    }
  }

  @Get('reset_password')
  @HttpCode(HttpStatus.OK)
  async resetPassword(@Query('token') token: string,
                      @Query('password') password: string): Promise<ResponseObject> {
    await this.authService.resetPassword(token,password);

    return {
      data: {},
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.RESET_PASSWORD,
    }
  }
}
