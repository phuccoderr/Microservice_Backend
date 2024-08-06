import { Body, Controller, Get, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateCustomerDto } from './dto/create-customer.dto';
import { ResponseObject } from '../response/response-object.dto';
import { ProducerService } from '../kafka/producer.service';
import { VerifyPayLoad } from './dto/kafka-verify.dto';
import { LoginCustomerDto } from './dto/login-customer.dto';
import { JwtPayload } from './dto/jwt-payload.dto';
import { ConfigService } from '@nestjs/config';

@Controller('api/v1/customers/auth')
export class AuthController {
  constructor(private readonly authService: AuthService,
              private readonly producerService: ProducerService,
              private readonly configService: ConfigService,) {}

  @Post('login')
  async login(@Body() loginCustomerDto: LoginCustomerDto): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.login(loginCustomerDto)

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Login Successfully!',
    };
  }

  @Post('refresh')
  async refreshToken(@Query('token') token: string): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.refreshToken(token);

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Refresh Token Successfully!',
    };
  }

  @Post('logout')
  async logout(@Query('token') token: string): Promise<ResponseObject> {
    const result: string = await this.authService.logout(token);

    return {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };
  }

  @Post('register')
  async register(@Body() createCustomerDto: CreateCustomerDto, ): Promise<ResponseObject> {

    const user = await this.authService.register(createCustomerDto);

    const url = this.configService.get("URI_VERIFY_REGISTER");

    const payload: VerifyPayLoad = {
      email: user.email,
      url: `${url}?token=${user.verification_code}`,
    }

   this.producerService.produce("customer-verify-events-topics", {
     value: JSON.stringify(payload),
   })

    return {
      data: url,
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

  @Post('forgot_password')
  async forgotPassword(@Query('email') email: string): Promise<ResponseObject> {
    const customer = await this.authService.forgotPassword(email);

    const url = this.configService.get("URI_VERIFY_PASSWORD");

    const payload: VerifyPayLoad = {
      email: email,
      url: `${url}?token=${customer.reset_password_token}`,
    }

    this.producerService.produce("customer-forgot-password-events-topics", {
      value: JSON.stringify(payload),
    })

    return {
      data: url,
      status: HttpStatus.OK,
      message: 'Verify forgot password successfully',
    }
  }

  @Post('reset_password')
  async resetPassword(@Query('token') token: string,
                      @Query('password') password: string): Promise<ResponseObject> {
    await this.authService.resetPassword(token,password);

    return {
      data: {},
      status: HttpStatus.OK,
      message: 'Reset password successfully',
    }
  }
}
