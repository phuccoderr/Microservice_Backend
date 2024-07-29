import { Injectable, UnauthorizedException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, Types } from 'mongoose';
import { RefreshToken } from './models/refresh-token.schema';
import { v4 as uuidv4 } from 'uuid';

@Injectable()
export class RefreshTokenRepository {
  constructor(
    @InjectModel(RefreshToken.name)
    private readonly refreshTokenModel: Model<RefreshToken>,
  ) {}

  async generateRefreshToken(userId: string): Promise<string> {
    const token = uuidv4();
    const expiresAt = new Date();
    expiresAt.setDate(expiresAt.getDate() + 7);

    new this.refreshTokenModel({
      token,
      userId,
      expiresAt: expiresAt,
      _id: new Types.ObjectId(),
    }).save();

    return token;
  }

  async findByToken(token: string): Promise<RefreshToken> {
    const refreshToken = await this.refreshTokenModel
      .findOne({ token })
      .lean<RefreshToken>(true);

    if (!refreshToken) {
      throw new UnauthorizedException("Can't find refresh token");
    }

    return refreshToken;
  }

  async findByUserIdAndDelete(userId: string): Promise<RefreshToken> {
    return await this.refreshTokenModel
      .findOneAndDelete({ userId })
      .lean<RefreshToken>(true);
  }

  async deleteByToken(refreshToken: string): Promise<void> {
    await this.refreshTokenModel.deleteOne({ token: refreshToken });
  }
}
