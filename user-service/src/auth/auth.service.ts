import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { LoginDto } from 'src/auth/dto/login.dto';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { RefreshTokenRepository } from 'src/auth/refresh-token.repository';
import { TokenPayLoad } from 'src/auth/interfaces/token-payload.interface';
import { JwtPayload } from 'src/auth/interfaces/jwt-payload.interface';

@Injectable()
export class AuthService {
  constructor(
    private readonly usersService: UsersService,
    private readonly jwtService: JwtService,
    private readonly refreshTokenRepository: RefreshTokenRepository,
  ) {}

  async login(request: LoginDto): Promise<JwtPayload> {
    const user = await this.usersService.verifyUser(request.email);

    const isMatch = await bcrypt.compare(request.password, user.password);
    if (!isMatch) {
      throw new UnauthorizedException('Password Not Match');
    }

    const tokenPayload: TokenPayLoad = {
      email: user.email,
      _id: user._id.toHexString(),
    };

    const access_token = this.jwtService.sign(tokenPayload);
    const refresh_token =
      await this.refreshTokenRepository.generateRefreshToken(
        user._id.toHexString(),
      );

    return { access_token, refresh_token };
  }
}
