import { Body, Controller, HttpStatus, Post, Query } from '@nestjs/common';
import { AuthService } from './auth.service';
import { LoginDto } from 'src/auth/dto/login.dto';
import { JwtPayload } from 'src/auth/interfaces/jwt-payload.interface';
import { ResponseObject } from 'src/response/response-object.dto';

@Controller('api/v1/')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  async login(@Body() loginDto: LoginDto): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.login(loginDto);

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Login Successfully!',
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

  @Post('refreshtoken')
  async refreshToken(@Query('token') token: string): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.refreshToken(token);

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: 'Refresh Token Successfully!',
    };
  }
}
