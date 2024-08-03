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

@Injectable()
export class AuthService {

  constructor(private readonly customersRepository: CustomersRepository,
              private readonly refreshTokenRepository: RefreshTokenRepository,
              private readonly jwtService: JwtService) {
  }

  async login(customerLoginDto: LoginCustomerDto): Promise<JwtPayload> {
    const customer = await this.customersRepository.findOne({email: customerLoginDto.email}, "");
    if (!customer) {
      throw new NotFoundException('Customer not found');
    }

    if (!customer.status) {
      throw new ForbiddenException('Customer not verify');
    }

    const isMatch = await bcrypt.compare(customerLoginDto.password, customer.password);
    if (!isMatch) {
      throw new UnauthorizedException('Password Not Match');
    }

    const tokenPayload: TokenPayload = {
      _id: customer._id.toHexString(),
      email: customerLoginDto.email,
    }

    await this.refreshTokenRepository.findByCustomerIdAndDelete(
      customer._id.toHexString(),
    );

    const accessToken = this.jwtService.sign(tokenPayload)
    const refreshToken = await this.refreshTokenRepository.generateRefreshToken(customer._id.toHexString());

    return {
      access_token: accessToken,
      refresh_token: refreshToken,
    };
  }

  async logout(token: string): Promise<string> {
    await this.refreshTokenRepository.findByTokenAndDelete(token)
    return 'Log out Successfully!';
  }

  async register(createCustomerDto: CreateCustomerDto): Promise<Customer> {
    try {
      return await this.customersRepository.create({
        address: '',
        authentication_type: AuthenticationType.DATABASE,
        phone_number: '',
        reset_password_token: '',
        status: false,
        verification_code: uuidv4(),
        ...createCustomerDto,
        password: await bcrypt.hash(createCustomerDto.password, 10)
      })
    } catch (error) {
      throw new UnprocessableEntityException(error.message);
    }
  }

  async verify (code: string): Promise<void> {

    const customer = await this.customersRepository.findByVerificationCode(code);
    if (!customer || customer.status) {
      throw new NotFoundException('Verify fail');
    }

    await this.customersRepository.findOneAndUpdate(
      {_id: customer._id},
      {verification_code: '', status: true});
  }

  async refreshToken (token: string): Promise<JwtPayload> {
    const customerRefreshToken  = await this.refreshTokenRepository.findByToken(token);
    if (!customerRefreshToken) {
      throw new NotFoundException("Can't find refresh token");
    }

    const customer = await this.customersRepository.findOne({_id: customerRefreshToken.customerId},"-password");
    if (!customer) {
      throw new NotFoundException("Can't find customer with refresh token");
    }

    await this.verifyToken(customerRefreshToken)

    const tokenPayload: TokenPayload = {
      _id: customer._id.toHexString(),
      email: customer.email,
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
      await this.refreshTokenRepository.findByTokenAndDelete(token);
      throw new UnauthorizedException('Token Expired');
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
