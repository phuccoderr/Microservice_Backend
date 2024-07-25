import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { LoginDto } from 'src/auth/dto/login.dto';
import { TokenPayLoad } from 'src/users/token-payload.interface';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private readonly usersService: UsersService,
    private readonly jwtService: JwtService,
  ) {}

  async login(request: LoginDto) {
    const user = await this.usersService.verifyUser(request.email);

    const isMatch = await bcrypt.compare(request.password, user.password);
    if (!isMatch) {
      throw new UnauthorizedException('Password Not Match');
    }

    const tokenPayload: TokenPayLoad = {
      email: user.email,
      _id: user._id.toHexString(),
    };
    return this.jwtService.sign(tokenPayload);
  }
}
