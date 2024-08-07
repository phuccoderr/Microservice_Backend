import { Injectable, NotFoundException, UnauthorizedException } from "@nestjs/common";
import { JwtService } from '@nestjs/jwt';
import { LoginDto } from './dto/login.dto';
import * as bcrypt from 'bcrypt';
import { RefreshTokenRepository } from './refresh-token.repository';
import { TokenPayLoad } from './dto/token-payload.dto';
import { JwtPayload } from './dto/jwt-payload.dto';
import { UserRefreshToken } from './models/refresh-token.schema';
import { UsersRepository } from '../users/users.repository';
import { DATABASE_CONST } from "@src/constants/db-constants";
import { AUTH_CONSTANTS } from "@src/constants/auth-constants";

@Injectable()
export class AuthService {
  constructor(
    private readonly jwtService: JwtService,
    private readonly refreshTokenRepository: RefreshTokenRepository,
    private readonly usersRepository: UsersRepository,
  ) {}

  async login(request: LoginDto): Promise<JwtPayload> {
    const user = await this.usersRepository.findOne({email: request.email}, "");
    if (!user) {
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }

    const isMatch = await bcrypt.compare(request.password, user.password);
    if (!isMatch) {
      throw new UnauthorizedException(AUTH_CONSTANTS.PASSWORD_NOT_MATCH);
    }

    const tokenPayload: TokenPayLoad = {
      email: user.email,
      _id: user._id.toHexString(),
      roles: user.roles,
    };

    await this.refreshTokenRepository.findOneAndDelete({_id: user._id.toHexString()});

    const accessToken = this.jwtService.sign(tokenPayload);
    const refreshToken = await this.refreshTokenRepository.generateRefreshToken(
      user._id.toHexString(),
    );

    return {
      access_token: accessToken,
      refresh_token: refreshToken,
    };
  }

  async logout(token: string): Promise<void> {
    await this.refreshTokenRepository.findOne({token},"");
  }

  async refreshToken(token: string): Promise<JwtPayload> {
    const refreshToken = await this.refreshTokenRepository.findOne({token},"");
    const user = await this.usersRepository.findOne(
      { _id: refreshToken.userId },
      '-password',
    );

    // check Expired
    await this.verifyToken(refreshToken);

    const tokenPayload: TokenPayLoad = {
      email: refreshToken.userId,
      _id: refreshToken.userId,
      roles: user.roles,
    };

    const accessToken = this.jwtService.sign(tokenPayload);

    return {
      access_token: accessToken,
      refresh_token: refreshToken.token,
    };
  }

  private async verifyToken(refreshToken: UserRefreshToken): Promise<string> {
    const { token, expiresAt } = refreshToken;

    if (new Date() > expiresAt) {
      await this.refreshTokenRepository.findOneAndDelete({token});
      throw new UnauthorizedException(AUTH_CONSTANTS.TOKEN_EXPIRED);
    }

    return token;
  }
}
