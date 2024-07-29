import { Injectable, Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { FilterQuery, Model } from 'mongoose';
import { AbstractRepository } from '../database/abstract.repository';
import { User } from './models/user.schema';

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
      .lean<User[]>(true);
  }

  async search(
    page: number,
    limit: number,
    sort: 'asc' | 'desc',
    filter: FilterQuery<User>,
  ) {
    return await this.model
      .find({ $or: [{ name: filter.name }, { email: filter.email }] })
      .sort({ name: sort })
      .skip((page - 1) * limit)
      .limit(limit)
      .select('-password')
      .lean<User[]>(true);
  }
}
