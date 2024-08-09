import {
  ForbiddenException,
  Injectable,
  NotFoundException,
  UnauthorizedException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CreateCustomerDto } from './dto/create-customer.dto';
import * as bcrypt from 'bcrypt';
import { AuthenticationType, Customer } from '../customers/models/customer.schema';
import { CustomersRepository } from '../customers/customers.repository';
import { v4 as uuidv4 } from 'uuid';
import { LoginCustomerDto } from './dto/login-customer.dto';
import { TokenPayload } from './dto/token-payload.dto';
import { RefreshTokenRepository } from './refresh-token.repository';
import { JwtService } from '@nestjs/jwt';
import { JwtPayload } from './dto/jwt-payload.dto';
import { CustomerRefreshToken } from './models/refresh-token.schema';
import { DATABASE_CONST } from '../constants/db-constants';
import { AUTH_CONSTANTS } from '../constants/auth-constants';

@Injectable()
export class AuthService {

  constructor(private readonly customersRepository: CustomersRepository,
              private readonly refreshTokenRepository: RefreshTokenRepository,
              private readonly jwtService: JwtService) {
  }

  async login(customerLoginDto: LoginCustomerDto): Promise<JwtPayload> {
    const customer = await this.customersRepository.findOne({email: customerLoginDto.email}, "");
    if (!customer) {
      throw new NotFoundException(DATABASE_CONST.NOTFOUND);
    }

    if (!customer.status) {
      throw new ForbiddenException(AUTH_CONSTANTS.VERIFY_ACCOUNT);
    }

    const isMatch = await bcrypt.compare(customerLoginDto.password, customer.password);
    if (!isMatch) {
      throw new UnauthorizedException(AUTH_CONSTANTS.PASSWORD_NOT_MATCH);
    }

    const tokenPayload: TokenPayload = {
      _id: customer._id.toHexString(),
      email: customer.email,
      roles: customer.roles
    }

    await this.refreshTokenRepository.findOneAndDelete(
      {customerId: customer._id.toHexString(),}
    );

    const accessToken = this.jwtService.sign(tokenPayload)
    const refreshToken = await this.refreshTokenRepository.generateRefreshToken(customer._id.toHexString());

    return {
      access_token: accessToken,
      refresh_token: refreshToken,
    };
  }

  async logout(token: string): Promise<void> {
    await this.refreshTokenRepository.findOneAndDelete({token: token})
  }

  async register(createCustomerDto: CreateCustomerDto): Promise<Customer> {
    try {
      return await this.customersRepository.create({
        authentication_type: AuthenticationType.DATABASE,
        status: false,
        verification_code: uuidv4(),
        ...createCustomerDto,
        password: await bcrypt.hash(createCustomerDto.password, 10)
      })
    } catch (error) {
      throw new UnprocessableEntityException(DATABASE_CONST.ALREADY);
    }
  }

  async verify (code: string): Promise<void> {

    const customer = await this.customersRepository.findByVerificationCode(code);
    if (!customer || customer.status) {
      throw new UnauthorizedException(AUTH_CONSTANTS.VERIFY_FAIL);
    }

    await this.customersRepository.findOneAndUpdate(
      {_id: customer._id},
      {verification_code: '', status: true});
  }

  async refreshToken (token: string): Promise<JwtPayload> {
    const customerRefreshToken  = await this.refreshTokenRepository.findOne({token}, "");
    if (!customerRefreshToken) {
      throw new NotFoundException(DATABASE_CONST.NOTFOUND + token);
    }

    const customer = await this.customersRepository.findOne({_id: customerRefreshToken.customerId},"-password");

    await this.verifyToken(customerRefreshToken)

    const tokenPayload: TokenPayload = {
      _id: customer._id.toHexString(),
      email: customer.email,
      roles: customer.roles
    }

    const accessToken = this.jwtService.sign(tokenPayload);

    return {
      access_token: accessToken,
      refresh_token: customerRefreshToken.token,
    }

  }
  private async verifyToken(refreshToken: CustomerRefreshToken): Promise<string> {
    const { token, expiresAt } = refreshToken;

    if (new Date() > expiresAt) {
      await this.refreshTokenRepository.findOneAndDelete({token});
      throw new UnauthorizedException(AUTH_CONSTANTS.TOKEN_EXPIRED);
    }

    return token;
  }

  async forgotPassword(email: string): Promise<Customer> {
    return await this.customersRepository.findOneAndUpdate(
      { email },
      { reset_password_token: uuidv4() });

  }

  async resetPassword(token: string,password: string): Promise<void> {
     await this.customersRepository.findOneAndUpdate(
      {reset_password_token: token},
      {reset_password_token: "", password: await bcrypt.hash(password, 10) });
  }
}
