import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { CustomerRefreshToken } from './models/refresh-token.schema';
import { Model, Types } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';

@Injectable()
export class RefreshTokenRepository {
  constructor(@InjectModel(CustomerRefreshToken.name) private readonly refreshTokenModel: Model<CustomerRefreshToken>) {}

  async generateRefreshToken(customerId: string): Promise<any> {
    const token = uuidv4();
    const expiresAt = new Date();
    expiresAt.setDate(expiresAt.getDate() + 7)

    await new this.refreshTokenModel({
      token,
      expiresAt,
      customerId,
      _id: new Types.ObjectId(),
    }).save();

    return token;
  }

  async findByCustomerIdAndDelete(customerId: string): Promise<CustomerRefreshToken> {
    return this.refreshTokenModel
      .findOneAndDelete({ customerId })
      .lean<CustomerRefreshToken>(true);
  }

  async findByTokenAndDelete(token: string): Promise<CustomerRefreshToken> {
    return this.refreshTokenModel.findOneAndDelete({ token: token }).lean<CustomerRefreshToken>(true);
  }

  async findByToken(token: string): Promise<CustomerRefreshToken> {
    return this.refreshTokenModel
      .findOneAndDelete({ token })
      .lean<CustomerRefreshToken>(true);
  }
}