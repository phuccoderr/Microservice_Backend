import { Body, Controller, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { LoginDto } from 'src/auth/dto/login.dto';
import { ResposneObject } from 'src/response/response-object.dto';
import { JwtPayload } from 'src/auth/interfaces/jwt-payload.interface';

@Controller('api/v1/')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  async login(@Body() loginDto: LoginDto) {
    const jwtPayload: JwtPayload = await this.authService.login(loginDto);

    const responseObject: ResposneObject = {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Login Successfully!',
    };

    return responseObject;
  }

  @Post('logout')
  async logout(@Query('token') token: string) {
    const result: string = await this.authService.logout(token);

    const responseObject: ResposneObject = {
      data: {},
      status: HttpStatus.OK,
      message: result,
    };

    return responseObject;
  }

  @Post('refreshtoken')
  async refreshToken(@Query('token') token: string) {
    const jwtPayload: JwtPayload = await this.authService.refreshToken(token);

    const responseObject: ResposneObject = {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Refresh Token Successfully!',
    };

    return responseObject;
  }
}
