import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { LoginDto } from 'src/auth/dto/login.dto';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { RefreshTokenRepository } from 'src/auth/refresh-token.repository';
import { TokenPayLoad } from 'src/auth/interfaces/token-payload.interface';
import { JwtPayload } from 'src/auth/interfaces/jwt-payload.interface';
import { RefreshToken } from 'src/auth/models/refresh-token.schema';
import { UsersRepository } from 'src/users/users.repository';

@Injectable()
export class AuthService {
  constructor(
    private readonly usersService: UsersService,
    private readonly jwtService: JwtService,
    private readonly refreshTokenRepository: RefreshTokenRepository,
    private readonly usersRepository: UsersRepository,
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
      roles: user.roles,
    };

    await this.refreshTokenRepository.findByUserIdAndDelete(
      user._id.toHexString(),
    );

    const accessToken = this.jwtService.sign(tokenPayload);
    const refreshToken = await this.refreshTokenRepository.generateRefreshToken(
      user._id.toHexString(),
    );

    const jwtPayload: JwtPayload = {
      access_token: accessToken,
      refresh_token: refreshToken,
    };

    return jwtPayload;
  }

  async logout(token: string): Promise<string> {
    const refreshToken = await this.refreshTokenRepository.findByToken(token);

    await this.refreshTokenRepository.deleteByToken(refreshToken.token);
    return 'Log out Successfully!';
  }

  async refreshToken(token: string): Promise<JwtPayload> {
    const refreshToken = await this.refreshTokenRepository.findByToken(token);
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

    const jwtPayload: JwtPayload = {
      access_token: accessToken,
      refresh_token: refreshToken.token,
    };

    return jwtPayload;
  }

  private async verifyToken(refreshToken: RefreshToken): Promise<string> {
    const { token, expiresAt } = refreshToken;

    if (new Date() > expiresAt) {
      this.refreshTokenRepository.deleteByToken(token);
      throw new UnauthorizedException('Token Expired');
    }

    return token;
  }
}
