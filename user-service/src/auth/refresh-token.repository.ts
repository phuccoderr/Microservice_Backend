import { Injectable, Logger } from "@nestjs/common";
import { InjectModel } from '@nestjs/mongoose';
import { Model, Types } from "mongoose";
import { UserRefreshToken } from "./models/refresh-token.schema";
import { v4 as uuidv4 } from 'uuid';
import { AbstractRepository } from "@src/database/abstract.repository";

@Injectable()
export class RefreshTokenRepository extends AbstractRepository<UserRefreshToken>{
  protected readonly logger: Logger;

  constructor(
    @InjectModel(UserRefreshToken.name)
    private readonly refreshTokenModel: Model<UserRefreshToken>,
  ) {
    super(refreshTokenModel)
  }

  async generateRefreshToken(userId: string): Promise<string> {
    const token = uuidv4();
    const expiresAt = new Date();
    expiresAt.setDate(expiresAt.getDate() + 7);

    await new this.refreshTokenModel({
      token,
      userId,
      expiresAt: expiresAt,
      _id: new Types.ObjectId(),
    }).save();

    return token;
  }


}
