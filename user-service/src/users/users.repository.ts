import { Injectable, Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { AbstractRepository } from 'src/database/abstract.repository';
import { User } from 'src/users/models/user.schema';

@Injectable()
export class UsersRepository extends AbstractRepository<User> {
  protected readonly logger = new Logger(UsersRepository.name);

  constructor(@InjectModel(User.name) userModel: Model<User>) {
    super(userModel);
  }

  async listByPage(page: number, limit: number, sort: 'asc' | 'desc') {
    return await this.model
      .find()
      .sort({ name: sort })
      .skip((page - 1) * limit)
      .limit(limit)
      .select('-password')
      .exec();
  }
}
