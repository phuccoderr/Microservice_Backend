import {
  Injectable,
  NotFoundException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UsersRepository } from './users.repository';
import * as bcrypt from 'bcrypt';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from './models/user.schema';
import { RequestPagination } from './dto/request-pagination.dto';

@Injectable()
export class UsersService {
  constructor(private readonly usersRepository: UsersRepository) {}

  async createUser(createUserDto: CreateUserDto): Promise<User> {
    try {
      return await this.usersRepository.create({
        ...createUserDto,
        password: await bcrypt.hash(createUserDto.password, 10),
      });
    } catch (error) {
      throw new UnprocessableEntityException('Email already exits');
    }
  }

  async getUsers(query: RequestPagination) {
    const { keyword, page, limit, sort } = query;

    if (keyword) {
      const filter = { email: keyword, name: keyword };
      return await this.usersRepository.search(page, limit, sort, filter);
    }

    const users = await this.usersRepository.listByPage(page, limit, sort);

    return users;
  }

  async getUser(_id: string) {
    try {
      const user = await this.usersRepository.findOne({ _id }, '-password');

      return user;
    } catch (error) {
      throw new NotFoundException('User not found');
    }
  }

  async updateUser(_id: string, updateUserDto: UpdateUserDto): Promise<User> {
    try {
      return await this.usersRepository.findOneAndUpdate(
        { _id },
        updateUserDto,
      );
    } catch (error) {
      throw new NotFoundException('User not found');
    }
  }

  async deleteUser(_id: string): Promise<string> {
    const userDelete = await this.usersRepository.findOneAndDelete({ _id });

    if (!userDelete) {
      throw new NotFoundException('User not found');
    }
    return 'Delete success';
  }

  async verifyUser(email: string) {
    return await this.usersRepository.findOne({ email }, '');
  }
}
