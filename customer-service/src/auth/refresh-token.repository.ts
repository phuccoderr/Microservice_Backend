import { Injectable, Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { CustomerRefreshToken } from './models/refresh-token.schema';
import { Model, Types } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';
import { AbstractRepository } from '../database/abstract.repository';


@Injectable()
export class RefreshTokenRepository extends AbstractRepository<CustomerRefreshToken>{

  protected readonly logger: Logger;
  constructor(@InjectModel(CustomerRefreshToken.name) private readonly refreshTokenModel: Model<CustomerRefreshToken>) {
    super(refreshTokenModel)
  }

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




}