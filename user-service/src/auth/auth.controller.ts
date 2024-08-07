import { Body, Controller, HttpCode, HttpStatus, Post, Query } from "@nestjs/common";
import { AuthService } from './auth.service';
import { LoginDto } from './dto/login.dto';
import { JwtPayload } from './dto/jwt-payload.dto';
import { ResponseObject } from '../response/response-object.dto';
import { AUTH_CONSTANTS } from "@src/constants/auth-constants";

@Controller('api/v1/users/auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  @HttpCode(HttpStatus.OK)
  async login(@Body() loginDto: LoginDto): Promise<ResponseObject> {
    const jwtPayload: JwtPayload = await this.authService.login(loginDto);

    return {
      data: jwtPayload,
      status: HttpStatus.OK,
      message: AUTH_CONSTANTS.LOGIN,
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
}
